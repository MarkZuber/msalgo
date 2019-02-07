package tokencache

type idTokenCacheItem struct {
	HomeAccountID  string `json:"home_account_id,omitempty"`
	Environment    string `json:"environment,omitempty"`
	RawClientInfo  string `json:"raw_client_info,omitempty"`
	CredentialType string `json:"credential_type,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	Secret         string `json:"secret,omitempty"`
	TenantID       string `json:"realm,omitempty"`
}
