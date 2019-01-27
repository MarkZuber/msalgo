package requests

type WsTrustResponse struct {
}

func (wsTrustResponse *WsTrustResponse) GetSAMLAssertion(endpoint *WsTrustEndpoint) (*SamlTokenInfo, error) {
	samlTokenInfo := &SamlTokenInfo{}
	return samlTokenInfo, nil
}
