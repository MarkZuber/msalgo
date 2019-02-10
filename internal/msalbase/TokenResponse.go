package msalbase

import (
	"encoding/json"
	"strings"
)

type tokenResponseJsonPayload struct {
	Error            string `json:"error"`
	SubError         string `json:"suberror"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	CorrelationID    string `json:"correlation_id"`
	Claims           string `json:"claims"`
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
}

type TokenResponse struct {
	errorVal         string
	subError         string
	errorDescription string
	errorCodes       []int
	correlationID    string
	claims           string
	accessToken      string
	refreshToken     string
	idToken          *IDToken
	grantedScopes    []string
}

func (tr *TokenResponse) IsAuthorizationPending() bool {
	return tr.errorVal == "authorization_pending"
}

func (tr *TokenResponse) GetAccessToken() string {
	return tr.accessToken
}

func (tr *TokenResponse) GetRefreshToken() string {
	return tr.refreshToken
}

func (tr *TokenResponse) GetIDToken() *IDToken {
	return tr.idToken
}

func (tr *TokenResponse) GetGrantedScopes() []string {
	return tr.grantedScopes
}

func (tr *TokenResponse) HasAccessToken() bool {
	return tr.accessToken != ""
}

func (tr *TokenResponse) HasRefreshToken() bool {
	return tr.refreshToken != ""
}

func (tr *TokenResponse) GetRawClientInfo() string {
	// todo:
	return ""
}

func CreateTokenResponse(authParameters *AuthParametersInternal, responseData string) (*TokenResponse, error) {
	payload := &tokenResponseJsonPayload{}
	var err = json.Unmarshal([]byte(responseData), payload)
	if err != nil {
		return nil, err
	}

	tokenResponse := &TokenResponse{
		errorVal:         payload.Error,
		subError:         payload.SubError,
		errorDescription: payload.ErrorDescription,
		errorCodes:       payload.ErrorCodes,
		correlationID:    payload.CorrelationID,
		claims:           payload.Claims,
		accessToken:      payload.AccessToken,
		refreshToken:     payload.RefreshToken}
	return tokenResponse, nil
}

func CreateTokenResponseFromParts(idToken *IDToken, accessToken *Credential, refreshToken *Credential) (*TokenResponse, error) {

	var idt *IDToken
	accessTokenSecret := ""
	refreshTokenSecret := ""
	grantedScopes := []string{}

	if idToken != nil {
		idt = idToken
	} else {
		idt = CreateIDToken("")
	}

	if accessToken != nil {
		accessTokenSecret = accessToken.Secret
		// todo: fill this in...
		// _expiresOn = TimeUtils::ToTimePoint(accessToken->GetExpiresOn());
		// _extendedExpiresOn = TimeUtils::ToTimePoint(accessToken->GetExtendedExpiresOn());
		grantedScopes = strings.Split(accessToken.Scopes, " ")
	}

	if refreshToken != nil {
		refreshTokenSecret = refreshToken.Secret
	}

	tokenResponse := &TokenResponse{idToken: idt, accessToken: accessTokenSecret, refreshToken: refreshTokenSecret, grantedScopes: grantedScopes}
	return tokenResponse, nil
}
