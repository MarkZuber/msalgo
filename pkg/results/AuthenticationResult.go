package results

// AuthenticationResult stuff
type AuthenticationResult struct {
	accessToken string
}

// CreateAuthenticationResult stuff
func CreateAuthenticationResult(accessToken string) *AuthenticationResult {
	ar := &AuthenticationResult{accessToken}
	return ar
}
