package requests

import "encoding/json"

type instanceDiscoveryResponse struct {
	TenantDiscoveryEndpoint string                      `json:"tenant_discovery_endpoint"`
	Metadata                []instanceDiscoveryMetadata `json:"metadata"`
}

func createInstanceDiscoveryResponse(responseData string) (*instanceDiscoveryResponse, error) {
	discoveryResponse := &instanceDiscoveryResponse{}
	var err = json.Unmarshal([]byte(responseData), discoveryResponse)
	if err != nil {
		return nil, err
	}
	return discoveryResponse, nil
}
