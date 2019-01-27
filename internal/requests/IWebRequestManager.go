package requests

// IWebRequestManager interface
type IWebRequestManager interface {
	GetUserRealm(authParameters *AuthParametersInternal) (*UserRealm, error)
	GetMex(federationMetadataURL string) (*WsTrustMexDocument, error)
	GetWsTrustResponse(authParameters *AuthParametersInternal, cloudAudienceURN string, endpoint *WsTrustEndpoint) (*WsTrustResponse, error)
	GetAccessTokenFromSamlGrant(authParameters *AuthParametersInternal, samlGrant *SamlTokenInfo) (*TokenResponse, error)
	GetAccessTokenFromUsernamePassword(authParameters *AuthParametersInternal) (*TokenResponse, error)
	GetAccessTokenFromAuthCode(authParameters *AuthParametersInternal, authCode string) (*TokenResponse, error)
	GetAccessTokenFromRefreshToken(authParameters *AuthParametersInternal, refreshToken string) (*TokenResponse, error)
	GetAccessTokenWithCertificate(authParameters *AuthParametersInternal, certificate *ClientCertificate) (*TokenResponse, error)
}
