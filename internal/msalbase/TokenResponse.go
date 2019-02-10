package msalbase

import (
	"encoding/json"
)

type TokenResponse struct {
	Error            string `json:"error"`
	SubError         string `json:"suberror"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	CorrelationID    string `json:"correlation_id"`
	Claims           string `json:"claims"`
	AccessToken      string `json:"access_token"`
}

func (tr *TokenResponse) IsAuthorizationPending() bool {
	return tr.Error == "authorization_pending"
}

func (tr *TokenResponse) GetAccessToken() string {
	return tr.AccessToken
}

func CreateTokenResponse(authParameters *AuthParametersInternal, responseData string) (*TokenResponse, error) {
	tokenResponse := &TokenResponse{}
	var err = json.Unmarshal([]byte(responseData), tokenResponse)
	if err != nil {
		return nil, err
	}
	return tokenResponse, nil
}
