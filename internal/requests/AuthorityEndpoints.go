package requests

type AuthorityEndpoints struct {
}

func (endpoints *AuthorityEndpoints) GetUserRealmEndpoint(username string) string {
	return "https://<environment>/common/UserRealm/<user>?api-version=1.0"
}
