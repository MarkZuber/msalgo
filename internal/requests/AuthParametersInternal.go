package requests

type AuthorizationType int

const (
	None AuthorizationType = iota
	UsernamePassword
	WindowsIntegratedAuth
	AuthCode
	Interactive
	Certificate
)

type AuthParametersInternal struct {
	authorityInfo     *AuthorityInfo
	correlationID     string
	endpoints         *AuthorityEndpoints
	clientID          string
	redirecturi       string
	accountid         string
	username          string
	password          string
	scopes            []string
	authorizationType AuthorizationType
}

func CreateAuthParametersInternal(clientID string) *AuthParametersInternal {
	p := &AuthParametersInternal{clientID: clientID}
	return p
}

func (ap *AuthParametersInternal) GetClientID() string {
	return ap.clientID
}

func (ap *AuthParametersInternal) GetCorrelationID() string {
	return ap.correlationID
}

func (ap *AuthParametersInternal) GetAuthorityEndpoints() *AuthorityEndpoints {
	return ap.endpoints
}

func (ap *AuthParametersInternal) GetUsername() string {
	return ap.username
}

func (ap *AuthParametersInternal) SetUsername(username string) {
	ap.username = username
}

func (ap *AuthParametersInternal) GetPassword() string {
	return ap.password
}

func (ap *AuthParametersInternal) SetPassword(password string) {
	ap.password = password
}

func (ap *AuthParametersInternal) GetRequestedScopes() []string {
	return ap.scopes
}

func (ap *AuthParametersInternal) GetAuthorityInfo() *AuthorityInfo {
	return ap.authorityInfo
}

func (ap *AuthParametersInternal) GetRedirectURI() string {
	return ap.redirecturi
}

func (ap *AuthParametersInternal) GetAuthorizationType() AuthorizationType {
	return ap.authorizationType
}

func (ap *AuthParametersInternal) SetAuthorizationType(authType AuthorizationType) {
	ap.authorizationType = authType
}
