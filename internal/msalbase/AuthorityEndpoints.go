package msalbase

import (
	"fmt"
	"net/url"
)

type AuthorityEndpoints struct {
}

func (endpoints *AuthorityEndpoints) GetUserRealmEndpoint(username string) string {
	return fmt.Sprintf("https://%s/common/UserRealm/%s?api-version=1.0", "login.microsoftonline.com", url.PathEscape(username))
}

func (endpoints *AuthorityEndpoints) GetTokenEndpoint() string {
	// todo: implement this stuff for real
	return "https://login.microsoftonline.com/organizations/oauth2/token"
}
