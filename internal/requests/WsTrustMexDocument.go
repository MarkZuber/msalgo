package requests

type WsTrustMexDocument struct {
	usernamePasswordEndpoint *WsTrustEndpoint
}

func (mexDoc *WsTrustMexDocument) GetWsTrustUsernamePasswordEndpoint() *WsTrustEndpoint {
	return mexDoc.usernamePasswordEndpoint
}
