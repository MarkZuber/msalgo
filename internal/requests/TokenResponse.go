package requests

type TokenResponse struct {
	accessToken string
}

func (tr *TokenResponse) GetAccessToken() string {
	return tr.accessToken
}
