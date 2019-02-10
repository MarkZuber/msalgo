package msalbase

type CredentialType int

const (
	CredentialTypeOauth2RefreshToken CredentialType = iota
	CredentialTypeOauth2AccessToken
	CredentialTypeOidcIdToken
	CredentialTypeOther
)

type Credential struct {
	HomeAccountID                  string
	Environment                    string
	RawClientInfo                  string
	CredentialType                 string // todo: credentialtype enum
	CredType                       CredentialType
	ClientID                       string
	Secret                         string
	Scopes                         string
	TenantID                       string
	ExpiresOnUnixTimestamp         string // todo: int64
	ExtendedExpiresOnUnixTimestamp string // todo: int64
	CachedAt                       string // todo: int64
	UserAssertionHash              string
	AdditionalFields               map[string]interface{}
}

func CreateCredentialRefreshToken(homeAccountID string, environment string, clientID string, cachedAt string, refreshToken string, additionalFieldsJson string) *Credential {
	c := &Credential{
		HomeAccountID: homeAccountID,
		Environment:   environment,
		ClientID:      clientID,
		CachedAt:      cachedAt,
		Secret:        refreshToken,
	}
	return c
}

func CreateCredentialAccessToken(homeAccountID string, environment string, realm string, clientID string, target string, cachedAt string, expiresOn string, extendedExpiresOn string, accessToken string, additionalFieldsJson string) *Credential {
	c := &Credential{
		HomeAccountID: homeAccountID,
		Environment:   environment,
		// todo: realm?
		ClientID:                       clientID,
		Scopes:                         target,
		CachedAt:                       cachedAt,
		ExpiresOnUnixTimestamp:         expiresOn,
		ExtendedExpiresOnUnixTimestamp: extendedExpiresOn,
		Secret:                         accessToken,
	}
	return c
}

func CreateCredentialIdToken(homeAccountID string, environment string, realm string, clientID string, cachedAt string, idTokenRaw string, additionalFieldsJson string) *Credential {
	c := &Credential{
		HomeAccountID: homeAccountID,
		Environment:   environment,
		// Realm:         realm,
		ClientID: clientID,
		CachedAt: cachedAt,
		Secret:   idTokenRaw,
	}
	return c
}
