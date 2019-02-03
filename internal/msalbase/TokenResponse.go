package msalbase

type TokenResponse struct {
	accessToken string
}

func (tr *TokenResponse) GetAccessToken() string {
	return tr.accessToken
}

func CreateTokenResponse(authParameters *AuthParametersInternal, responseData string) (*TokenResponse, error) {
	tr := &TokenResponse{}
	return tr, nil
}
