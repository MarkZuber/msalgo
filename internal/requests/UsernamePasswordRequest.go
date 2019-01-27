package requests

// UsernamePasswordRequest stuff
type UsernamePasswordRequest struct {
	webRequestManager IWebRequestManager
	authParameters    *AuthParametersInternal
}

// CreateUsernamePasswordRequest stuff
func CreateUsernamePasswordRequest(
	webRequestManager IWebRequestManager,
	authParameters *AuthParametersInternal) *UsernamePasswordRequest {
	req := &UsernamePasswordRequest{webRequestManager, authParameters}
	return req
}

// Execute stuff
func (req *UsernamePasswordRequest) Execute() (*TokenResponse, error) {
	_, err := req.webRequestManager.GetUserRealm(req.authParameters)
	if err != nil {
		return nil, err
	}

	return nil, nil
	// switch accountType := userRealm.GetAccountType(); accountType {
	// 	return req.webRequestManager.GetAccessTokenFromUsernamePassword(req.authParameters)
	// }
}
