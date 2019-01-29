package requests

type SamlAssertionType int

const (
	SamlV1 SamlAssertionType = iota
	SamlV2
)

type SamlTokenInfo struct {
	assertionType SamlAssertionType
	assertion     string
}

func CreateSamlTokenInfo(assertionType SamlAssertionType, assertion string) *SamlTokenInfo {
	tokenInfo := &SamlTokenInfo{assertionType, assertion}
	return tokenInfo
}
