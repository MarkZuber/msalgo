package requests

type tenantDiscoveryResponse struct {
	Error            string `json:"error"`
	SubError         string `json:"suberror"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	CorrelationID    string `json:"correlation_id"`
	Claims           string `json:"claims"`

	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	Issuer                string `json:"issuer"`
}

func (r *tenantDiscoveryResponse) HasAuthorizationEndpoint() bool {
	return len(r.AuthorizationEndpoint) > 0
}

func (r *tenantDiscoveryResponse) HasTokenEndpoint() bool {
	return len(r.TokenEndpoint) > 0
}

func (r *tenantDiscoveryResponse) HasIssuer() bool {
	return len(r.Issuer) > 0
}
