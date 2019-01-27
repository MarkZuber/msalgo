package requests

type AuthParametersInternal struct {
	authority   string
	endpoints   *AuthorityEndpoints
	clientID    string
	redirecturi string
	accountid   string
	username    string
	password    string
}

func CreateAuthParametersInternal(clientID string) *AuthParametersInternal {
	p := &AuthParametersInternal{clientID: clientID}
	return p
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
