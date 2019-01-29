package requests

import (
	"errors"
	"log"
)

// WebRequestManager stuff
type WebRequestManager struct {
	httpManager *HTTPManager
}

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
	return nil, nil
}

func (wrm *WebRequestManager) GetAccessTokenFromUsernamePassword(authParameters *AuthParametersInternal) (*TokenResponse, error) {
	return nil, nil
}

func (wrm *WebRequestManager) GetAccessTokenFromAuthCode(authParameters *AuthParametersInternal, authCode string) (*TokenResponse, error) {
	return nil, nil
}

func (wrm *WebRequestManager) GetAccessTokenFromRefreshToken(authParameters *AuthParametersInternal, refreshToken string) (*TokenResponse, error) {
	return nil, nil
}

func (wrm *WebRequestManager) GetAccessTokenWithCertificate(authParameters *AuthParametersInternal, certificate *ClientCertificate) (*TokenResponse, error) {
	return nil, nil
}

func (wrm *WebRequestManager) getAadHeaders(authParameters *AuthParametersInternal) map[string]string {
	headers := make(map[string]string)
	return headers
}
