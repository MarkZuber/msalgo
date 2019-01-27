package requests

type AuthParametersInternal struct {
	authority   string
	endpoints   *AuthorityEndpoints
	clientid    string
	redirecturi string
	accountid   string
	username    string
	password    string
}

func (ap *AuthParametersInternal) GetAuthorityEndpoints() *AuthorityEndpoints {
	return ap.endpoints
}

func (ap *AuthParametersInternal) GetUserName() string {
	return ap.username
}
