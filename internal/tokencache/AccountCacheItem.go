package tokencache

type accountCacheItem struct {
	HomeAccountID     string `json:"home_account_id,omitempty"`
	Environment       string `json:"environment,omitempty"`
	RawClientInfo     string `json:"raw_client_info,omitempty"`
	TenantID          string `json:"realm,omitempty"`
	PreferredUsername string `json:"username,omitempty"`
	Name              string `json:"name,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	FamilyName        string `json:"family_name,omitempty"`
	LocalAccountID    string `json:"local_account_id,omitempty"`
	AuthorityType     string `json:"authority_type,omitempty"`
}
