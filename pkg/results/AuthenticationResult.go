package results

import "github.com/markzuber/msalgo/internal/requests"

// AuthenticationResult stuff
type AuthenticationResult struct {
	tokenResponse *requests.TokenResponse
}

// CreateAuthenticationResult stuff
func CreateAuthenticationResult(tokenResponse *requests.TokenResponse) *AuthenticationResult {
	ar := &AuthenticationResult{tokenResponse}
	return ar
}

func (ar *AuthenticationResult) GetAccessToken() string {
	return ar.tokenResponse.GetAccessToken()
}
