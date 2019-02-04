package requests

type instanceDiscoveryResponse struct {
	TenantDiscoveryEndpoint string                      `json:"tenant_discovery_endpoint"`
	Metadata                []instanceDiscoveryMetadata `json:"metadata"`
}
