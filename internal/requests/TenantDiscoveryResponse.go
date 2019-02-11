package requests

import (
	"encoding/json"

	"github.com/markzuber/msalgo/internal/msalbase"
)

type tenantDiscoveryResponse struct {
	BaseResponse *msalbase.OAuthResponseBase

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

func createTenantDiscoveryResponse(responseCode int, responseData string) (*tenantDiscoveryResponse, error) {
	baseResponse, err := msalbase.CreateOAuthResponseBase(responseCode, responseData)
	if err != nil {
		return nil, err
	}

	discoveryResponse := &tenantDiscoveryResponse{}
	err = json.Unmarshal([]byte(responseData), discoveryResponse)
	if err != nil {
		return nil, err
	}

	discoveryResponse.BaseResponse = baseResponse

	return discoveryResponse, nil
}
