package requests

import (
	"bytes"
	"errors"
	"log"
	"net/url"
)

// WebRequestManager stuff
type WebRequestManager struct {
	httpManager *HTTPManager
}

type ContentType int

const (
	SoapXmlUtf8 ContentType = iota
	UrlEncodedUtf8
)

func CreateWebRequestManager(httpManager *HTTPManager) IWebRequestManager {
	m := &WebRequestManager{httpManager}
	return m
}

// GetUserRealm stuff
func (wrm *WebRequestManager) GetUserRealm(authParameters *AuthParametersInternal) (*UserRealm, error) {
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

	return CreateUserRealm(httpManagerResponse.GetResponseData())
}

func (wrm *WebRequestManager) GetMex(federationMetadataURL string) (*WsTrustMexDocument, error) {
	httpManagerResponse, err := wrm.httpManager.Get(federationMetadataURL, nil)
	if err != nil {
		return nil, err
	}

	if httpManagerResponse.GetResponseCode() != 200 {
		return nil, errors.New("invalid response code") // todo: need error struct here
	}

	return CreateWsTrustMexDocument(httpManagerResponse.GetResponseData())
}

func (wrm *WebRequestManager) GetWsTrustResponse(authParameters *AuthParametersInternal, cloudAudienceURN string, endpoint *WsTrustEndpoint) (*WsTrustResponse, error) {
	return nil, nil
}

func (wrm *WebRequestManager) GetAccessTokenFromSamlGrant(authParameters *AuthParametersInternal, samlGrant *SamlTokenInfo) (*TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type": "password",
		"username":   authParameters.GetUsername(),
		"password":   authParameters.GetPassword(),
	}

	switch samlGrant.GetAssertionType() {
	case SamlV1:
		decodedQueryParams["grant_type"] = "urn:ietf:params:oauth:grant-type:saml1_1-bearer"
		break
	case SamlV2:
		decodedQueryParams["grant_type"] = "urn:ietf:params:oauth:grant-type:saml2-bearer"
		break
	default:
		return nil, errors.New("GetAccessTokenFromSamlGrant returned unknown saml assertion type")
		// MSAL_THROW(
		//     UNTAGGED,
		//     Status::Unexpected,
		//     "GetAccessTokenFromSamlGrant returned unknown saml assertion type: '%d'",
		//     static_cast<int32_t>(samlGrant->GetAssertionType()));
	}
	// todo: decodedQueryParams["assertion"] = StringUtils::Base64RFCEncodePadded(samlGrant->GetAssertion());
	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetAccessTokenFromUsernamePassword(authParameters *AuthParametersInternal) (*TokenResponse, error) {

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

func addClientIdQueryParam(queryParams map[string]string, authParameters *AuthParametersInternal) {
	queryParams["client_id"] = authParameters.GetClientID()
}

func joinScopes(scopes []string) string {
	var buffer bytes.Buffer
	for _, scope := range scopes {
		buffer.WriteString(scope)
	}
	return buffer.String()
}

func addScopeQueryParam(queryParams map[string]string, authParameters *AuthParametersInternal) {
	// MSAL_DEBUG("Adding scopes 'openid', 'offline_access', 'profile'");
	requestedScopes := authParameters.GetRequestedScopes()

	// openid equired to get an id token
	// offline_access required to get a refresh token
	// profile required to get the client_info field back
	requestedScopes = append(requestedScopes, "openid", "offline_access", "profile")
	queryParams["scope"] = joinScopes(requestedScopes)
}

func addClientInfoQueryParam(queryParams map[string]string) {
	queryParams["client_info"] = "1"
}

func addRedirectUriQueryParam(queryParams map[string]string, authParameters *AuthParametersInternal) {
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

func getAadHeaders(authParameters *AuthParametersInternal) map[string]string {
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

	first := true
	for k, v := range queryParameters {
		if !first {
			buffer.WriteString("&")
			first = false
		}
		buffer.WriteString(url.QueryEscape(k))
		buffer.WriteString("=")
		buffer.WriteString(url.QueryEscape(v))
	}

	return buffer.String()
}

func (wrm *WebRequestManager) exchangeGrantForToken(authParameters *AuthParametersInternal, queryParams map[string]string) (*TokenResponse, error) {
	headers := getAadHeaders(authParameters)
	addContentTypeHeader(headers, UrlEncodedUtf8)

	response, err := wrm.httpManager.Post(
		authParameters.GetAuthorityInfo().GetTokenEndpoint(), encodeQueryParameters(queryParams), headers)
	if err != nil {
		return nil, err
	}
	return CreateTokenResponse(authParameters, response.GetResponseData())
}

func (wrm *WebRequestManager) GetAccessTokenFromAuthCode(authParameters *AuthParametersInternal, authCode string) (*TokenResponse, error) {
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

func (wrm *WebRequestManager) GetAccessTokenFromRefreshToken(authParameters *AuthParametersInternal, refreshToken string) (*TokenResponse, error) {
	decodedQueryParams := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
	}

	addClientIdQueryParam(decodedQueryParams, authParameters)
	addScopeQueryParam(decodedQueryParams, authParameters)
	addClientInfoQueryParam(decodedQueryParams)

	return wrm.exchangeGrantForToken(authParameters, decodedQueryParams)
}

func (wrm *WebRequestManager) GetAccessTokenWithCertificate(authParameters *AuthParametersInternal, certificate *ClientCertificate) (*TokenResponse, error) {

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

func (wrm *WebRequestManager) getAadHeaders(authParameters *AuthParametersInternal) map[string]string {
	headers := make(map[string]string)
	return headers
}
