package requests

import (
	"bytes"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/internal/wstrust"
	"github.com/markzuber/msalgo/pkg/contracts"
)

// WebRequestManager stuff
type WebRequestManager struct {
	httpManager *msalbase.HTTPManager
}

func isErrorAuthorizationPending(err error) bool {
	// todo: implement me!
	return false
}

type ContentType int

const (
	SoapXmlUtf8 ContentType = iota
	UrlEncodedUtf8
)

func CreateWebRequestManager(httpManager *msalbase.HTTPManager) IWebRequestManager {
	m := &WebRequestManager{httpManager}
	return m
}

// GetUserRealm stuff
func (wrm *WebRequestManager) GetUserRealm(authParameters *msalbase.AuthParametersInternal) (*msalbase.UserRealm, error) {
	log.Println("getuserrealm entered")
	url := authParameters.GetAuthorityEndpoints().GetUserRealmEndpoint(authParameters.GetUsername())

	log.Println("user realm endpoint: " + url)
	httpManagerResponse, err := wrm.httpManager.Get(url, wrm.getAadHeaders(authParameters))
	if err != nil {
		return nil, err
	}

	if httpManagerResponse.GetResponseCode() != 200 {
		return nil, errors.New("invalid response code") // todo: need error struct here
	}

	return msalbase.CreateUserRealm(httpManagerResponse.GetResponseData())
}

func (wrm *WebRequestManager) GetMex(federationMetadataURL string) (*wstrust.WsTrustMexDocument, error) {
	httpManagerResponse, err := wrm.httpManager.Get(federationMetadataURL, nil)
	if err != nil {
		return nil, err
	}

	if httpManagerResponse.GetResponseCode() != 200 {
		return nil, errors.New("invalid response code") // todo: need error struct here
	}

	return wstrust.CreateWsTrustMexDocument(httpManagerResponse.GetResponseData())
}

func (wrm *WebRequestManager) GetWsTrustResponse(authParameters *msalbase.AuthParametersInternal, cloudAudienceURN string, endpoint *wstrust.WsTrustEndpoint) (*wstrust.WsTrustResponse, error) {
	var wsTrustRequestMessage string
	var err error

	switch authParameters.GetAuthorizationType() {
	case msalbase.WindowsIntegratedAuth:
		wsTrustRequestMessage, err = endpoint.BuildTokenRequestMessageWIA(cloudAudienceURN)
	case msalbase.UsernamePassword:
		wsTrustRequestMessage, err = endpoint.BuildTokenRequestMessageUsernamePassword(
			cloudAudienceURN, authParameters.GetUsername(), authParameters.GetPassword())
	default:
		log.Println("unknown auth type!")
		err = errors.New("Unknown auth type")
	}

	if err != nil {
		return nil, err
	}

	var soapAction string

	// todo: make consts out of these strings
	if endpoint.GetVersion() == wstrust.Trust2005 {
		soapAction = "http://schemas.xmlsoap.org/ws/2005/02/trust/RST/Issue"
	} else {
		soapAction = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/RST/Issue"
	}

	headers := map[string]string{
		"SOAPAction": soapAction,
	}

	addContentTypeHeader(headers, SoapXmlUtf8)

	log.Println("calling POST for wstrustresponse")
	log.Println(endpoint.GetURL())
	log.Println(wsTrustRequestMessage)

	response, err := wrm.httpManager.Post(endpoint.GetURL(), wsTrustRequestMessage, headers)
	if err != nil {
		return nil, err
	}

	log.Println(response.GetResponseData())

	return wstrust.CreateWsTrustResponse(response.GetResponseData()), nil
}

func (wrm *WebRequestManager) GetAccessTokenFromSamlGrant(authParameters *msalbase.AuthParametersInternal, samlGrant *wstrust.SamlTokenInfo) (*msalbase.TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type": "password",
		"username":   authParameters.GetUsername(),
		"password":   authParameters.GetPassword(),
	}

	switch samlGrant.GetAssertionType() {
	case wstrust.SamlV1:
		decodedQueryParams["grant_type"] = "urn:ietf:params:oauth:grant-type:saml1_1-bearer"
		break
	case wstrust.SamlV2:
		decodedQueryParams["grant_type"] = "urn:ietf:params:oauth:grant-type:saml2-bearer"
		break
	default:
		return nil, errors.New("GetAccessTokenFromSamlGrant returned unknown saml assertion type: " + string(samlGrant.GetAssertionType()))
	}
	// todo: decodedQueryParams["assertion"] = StringUtils::Base64RFCEncodePadded(samlGrant->GetAssertion());
	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetAccessTokenFromUsernamePassword(authParameters *msalbase.AuthParametersInternal) (*msalbase.TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type": "password",
		"username":   authParameters.GetUsername(),
		"password":   authParameters.GetPassword(),
	}

	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetDeviceCodeResult(authParameters *msalbase.AuthParametersInternal) (*contracts.DeviceCodeResult, error) {
	decodedQueryParams := map[string]string{}

	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)

	deviceCodeEndpoint := strings.Replace(authParameters.GetAuthorityEndpoints().GetTokenEndpoint(), "token", "devicecode", -1)

	headers := getAadHeaders(authParameters)
	addContentTypeHeader(headers, UrlEncodedUtf8)

	response, err := wrm.httpManager.Post(
		deviceCodeEndpoint, encodeQueryParameters(decodedQueryParams), headers)
	if err != nil {
		return nil, err
	}
	dcResponse, err := createDeviceCodeResponse(response.GetResponseData())
	if err != nil {
		return nil, err
	}

	return dcResponse.toDeviceCodeResult(authParameters.GetClientID(), authParameters.GetScopes()), nil
}

func (wrm *WebRequestManager) GetAccessTokenFromDeviceCodeResult(authParameters *msalbase.AuthParametersInternal, deviceCodeResult *contracts.DeviceCodeResult) (*msalbase.TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type":  "device_code",
		"device_code": deviceCodeResult.GetDeviceCode(),
	}

	addClientIdQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)
	addScopeQueryParam(decodedQueryParams, authParameters)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func addClientIdQueryParam(queryParams map[string]string, authParameters *msalbase.AuthParametersInternal) {
	queryParams["client_id"] = authParameters.GetClientID()
}

func joinScopes(scopes []string) string {
	return strings.Join(scopes[:], " ")
}

func addScopeQueryParam(queryParams map[string]string, authParameters *msalbase.AuthParametersInternal) {
	log.Println("Adding scopes 'openid', 'offline_access', 'profile'")
	requestedScopes := authParameters.GetScopes()

	// openid equired to get an id token
	// offline_access required to get a refresh token
	// profile required to get the client_info field back
	requestedScopes = append(requestedScopes, "openid", "offline_access", "profile")
	queryParams["scope"] = joinScopes(requestedScopes)
}

func addClientInfoQueryParam(queryParams map[string]string) {
	queryParams["client_info"] = "1"
}

func addRedirectUriQueryParam(queryParams map[string]string, authParameters *msalbase.AuthParametersInternal) {
	queryParams["redirect_uri"] = authParameters.GetRedirectURI()
}

func addContentTypeHeader(headers map[string]string, contentType ContentType) {
	contentTypeKey := "Content-Type"
	switch contentType {
	case SoapXmlUtf8:
		headers[contentTypeKey] = "application/soap+xml; charset=utf-8"
		return

	case UrlEncodedUtf8:
		headers[contentTypeKey] = "application/x-www-form-urlencoded; charset=utf-8"
		return
	}
}

func getAadHeaders(authParameters *msalbase.AuthParametersInternal) map[string]string {
	headers := make(map[string]string)

	// headers["x-client-SKU"] = FormatUtils::FormatString("MSAL.golang.%s", systemInfo->GetName());
	// headers["x-client-OS"] = systemInfo->GetVersion();
	// headers["x-client-Ver"] = MSAL_VERSION_FROM_CMAKE;
	headers["client-request-id"] = authParameters.GetCorrelationID()
	headers["return-client-request-id"] = "false"
	return headers
}

func encodeQueryParameters(queryParameters map[string]string) string {
	var buffer bytes.Buffer

	log.Println("encodeQueryParameters...")

	first := true
	for k, v := range queryParameters {
		if !first {
			buffer.WriteString("&")
		}
		first = false
		buffer.WriteString(url.QueryEscape(k))
		buffer.WriteString("=")
		buffer.WriteString(url.QueryEscape(v))
	}

	result := buffer.String()
	log.Println(result)
	return result
}

func (wrm *WebRequestManager) exchangeGrantForToken(authParameters *msalbase.AuthParametersInternal, queryParams map[string]string) (*msalbase.TokenResponse, error) {
	headers := getAadHeaders(authParameters)
	addContentTypeHeader(headers, UrlEncodedUtf8)

	response, err := wrm.httpManager.Post(authParameters.GetAuthorityEndpoints().GetTokenEndpoint(), encodeQueryParameters(queryParams), headers)
	if err != nil {
		return nil, err
	}
	return msalbase.CreateTokenResponse(authParameters, response.GetResponseData())
}

func (wrm *WebRequestManager) GetAccessTokenFromAuthCode(authParameters *msalbase.AuthParametersInternal, authCode string) (*msalbase.TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type": "authorization_code",
		"code":       authCode,
	}

	addRedirectUriQueryParam(decodedQueryParams, authParameters)
	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetAccessTokenFromRefreshToken(authParameters *msalbase.AuthParametersInternal, refreshToken string) (*msalbase.TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetAccessTokenWithCertificate(authParameters *msalbase.AuthParametersInternal, certificate *msalbase.ClientCertificate) (*msalbase.TokenResponse, error) {

	assertion := "GetClientCertForAudience()" // todo: string assertion = GetClientCertificateAssertionForAudience(authParameters, certificate);

	decodedQueryParams := map[string]string{
		"grant_type":            "client_credentials",
		"client_assertion_type": "urn:ietf:params:oauth:client-assertion-type:jwt-bearer",
		"client_assertion":      assertion,
	}

	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)
	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) getAadHeaders(authParameters *msalbase.AuthParametersInternal) map[string]string {
	headers := make(map[string]string)
	return headers
}

func (wrm *WebRequestManager) GetAadinstanceDiscoveryResponse(authorityInfo *msalbase.AuthorityInfo) (*instanceDiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}

func (wrm *WebRequestManager) GetTenantDiscoveryResponse(openIdConfigurationEndpoint string) (*tenantDiscoveryResponse, error) {
	return nil, errors.New("not implemented")
}
