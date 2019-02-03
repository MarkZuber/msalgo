package msalbase

type AuthorityInfo struct {
	authorityURI string
}

func (ai *AuthorityInfo) GetTokenEndpoint() string {
	return "todo: implement returning the token endpoint"
}
