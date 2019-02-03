package msalgo

import (
	"github.com/markzuber/msalgo/internal/msalbase"
)

// AuthenticationResult contains the results of one token acquisition operation in PublicClientApplication
// or ConfidentialClientApplication. For details see https://aka.ms/msal-net-authenticationresult
type AuthenticationResult struct {
	tokenResponse *msalbase.TokenResponse
}

// CreateAuthenticationResult creates and AuthenticationResult.  This should only be called from internal code.
func createAuthenticationResult(tokenResponse *msalbase.TokenResponse) *AuthenticationResult {
	ar := &AuthenticationResult{tokenResponse}
	return ar
}

// GetAccessToken retrieves the access token string from the result.
func (ar *AuthenticationResult) GetAccessToken() string {
	return ar.tokenResponse.GetAccessToken()
}
