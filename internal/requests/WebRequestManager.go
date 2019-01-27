package requests

import "errors"

// WebRequestManager stuff
type WebRequestManager struct {
	httpManager *HTTPManager
}

// GetUserRealm stuff
func (wrm *WebRequestManager) GetUserRealm(authParameters *AuthParametersInternal) (*UserRealm, error) {
	url := authParameters.GetAuthorityEndpoints().GetUserRealmEndpoint(authParameters.GetUserName())
	httpManagerResponse, err := wrm.httpManager.Get(url, wrm.getAadHeaders(authParameters))
	if err != nil {
		return nil, err
	}

	if httpManagerResponse.GetResponseCode() != 200 {
		return nil, errors.New("invalid response code") // todo: need error struct here
	}

	return CreateUserRealm(httpManagerResponse.GetResponseData())
}

func (wrm *WebRequestManager) getAadHeaders(authParameters *AuthParametersInternal) map[string]string {
	headers := make(map[string]string)
	return headers
}
