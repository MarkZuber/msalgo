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
	return "todo: implement returning the token endpoint"
}
