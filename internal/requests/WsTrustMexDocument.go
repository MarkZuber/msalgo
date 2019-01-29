package requests

import "log"

type WsTrustMexDocument struct {
	usernamePasswordEndpoint *WsTrustEndpoint
}

func CreateWsTrustMexDocument(responseData string) (*WsTrustMexDocument, error) {
	log.Println(responseData)
	return nil, nil
}

func (mexDoc *WsTrustMexDocument) GetWsTrustUsernamePasswordEndpoint() *WsTrustEndpoint {
	return mexDoc.usernamePasswordEndpoint
}
