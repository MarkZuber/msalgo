package tokencache

import "encoding/json"

type accessTokenCacheItem struct {
	HomeAccountID                  string `json:"home_account_id,omitempty"`
	Environment                    string `json:"environment,omitempty"`
	RawClientInfo                  string `json:"raw_client_info,omitempty"`
	CredentialType                 string `json:"credential_type,omitempty"`
	ClientID                       string `json:"client_id,omitempty"`
	Secret                         string `json:"secret,omitempty"`
	Scopes                         string `json:"target,omitempty"`
	TenantID                       string `json:"tenant_id,omitempty"`
	ExpiresOnUnixTimestamp         string `json:"expires_on,omitempty"`
	ExtendedExpiresOnUnixTimestamp string `json:"extended_expires_on,omitempty"`
	CachedAt                       string `json:"cached_at,omitempty"`
	UserAssertionHash              string `json:"user_assertion_hash,omitempty"`
	AdditionalFields               map[string]interface{}
}

func extractExistingOrEmptyString(j map[string]interface{}, key string) string {
	if val, ok := j[key]; ok {
		if str, ok := val.(string); ok {
			delete(j, key)
			return str
		}
	}
	return ""
}

func (s *accessTokenCacheItem) populateFromJSONMap(j map[string]interface{}) error {
	s.HomeAccountID = extractExistingOrEmptyString(j, "home_account_id")
	s.AdditionalFields = j
	return nil
}

func (s *accessTokenCacheItem) UnmarshalJSON(b []byte) error {
	j := make(map[string]interface{})
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	return s.populateFromJSONMap(j)
}

func (s *accessTokenCacheItem) toJSONMap() map[string]interface{} {
	j := make(map[string]interface{})
	for k, v := range s.AdditionalFields {
		j[k] = v
	}

	j["home_account_id"] = s.HomeAccountID

	return j
}

func (s *accessTokenCacheItem) MarshalJSON() ([]byte, error) {
	j := s.toJSONMap()
	return json.Marshal(j)
}
