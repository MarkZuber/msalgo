package msalbase

type IDToken struct {
	PreferredUsername string `json:"preferred_username,omitempty"`
	GivenName         string `json:"given_name,omitempty"`
	FamilyName        string `json:"family_name,omitempty"`
	MiddleName        string `json:"middle_name,omitempty"`
	Name              string `json:"name,omitempty"`
	Oid               string `json:"oid,omitempty"`
	TenantID          string `json:"tenant_id,omitempty"`
	Subject           string `json:"subject,omitempty"`
	UPN               string `json:"upn,omitempty"`
	Email             string `json:"email,omitempty"`
	AlternativeID     string `json:"alternative_id,omitempty"`
}

func CreateIDToken(jwt string) *IDToken {
	// todo: unmarshal the json...
	return nil
}

func (t *IDToken) GetRaw() string {
	return "" // todo:
}

func (t *IDToken) IsEmpty() bool {
	return true // todo:
}

func (t *IDToken) GetPreferredUsername() string {
	return t.PreferredUsername
}

func (t *IDToken) GetGivenName() string {
	return t.GivenName
}

func (t *IDToken) GetFamilyName() string {
	return t.FamilyName
}

func (t *IDToken) GetMiddleName() string {
	return t.MiddleName
}

func (t *IDToken) GetName() string {
	return t.Name
}

func (t *IDToken) GetAlternativeId() string {
	return t.AlternativeID
}
