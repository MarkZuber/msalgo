package msalbase

import (
	"encoding/json"
	"log"
)

type TokenResponse struct {
	Error            string `json:"error"`
	SubError         string `json:"suberror"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	CorrelationID    string `json:"correlation_id"`
	Claims           string `json:"claims"`

	accessToken string
}

func (tr *TokenResponse) IsAuthorizationPending() bool {
	log.Println(tr.Error)
	return tr.Error == "authorization_pending"
}

func (tr *TokenResponse) GetAccessToken() string {
	return tr.accessToken
}

func CreateTokenResponse(authParameters *AuthParametersInternal, responseData string) (*TokenResponse, error) {
	tokenResponse := &TokenResponse{}
	var err = json.Unmarshal([]byte(responseData), tokenResponse)
	if err != nil {
		return nil, err
	}
	return tokenResponse, nil
}
