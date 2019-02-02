package requests

import (
	"errors"
	"log"
)

// UsernamePasswordRequest stuff
type UsernamePasswordRequest struct {
	webRequestManager IWebRequestManager
	authParameters    *AuthParametersInternal
}

// CreateUsernamePasswordRequest stuff
func CreateUsernamePasswordRequest(
	webRequestManager IWebRequestManager,
	authParameters *AuthParametersInternal) *UsernamePasswordRequest {
	req := &UsernamePasswordRequest{webRequestManager, authParameters}
	return req
}

// Execute stuff
func (req *UsernamePasswordRequest) Execute() (*TokenResponse, error) {
	userRealm, err := req.webRequestManager.GetUserRealm(req.authParameters)
	if err != nil {
		return nil, err
	}

	log.Println("got user realm")

	switch accountType := userRealm.GetAccountType(); accountType {
	case Federated:
		log.Println("FEDERATED")
		if mexDoc, err := req.webRequestManager.GetMex(userRealm.GetFederationMetadataURL()); err == nil {
			wsTrustEndpoint := mexDoc.GetWsTrustUsernamePasswordEndpoint()
			if wsTrustResponse, err := req.webRequestManager.GetWsTrustResponse(req.authParameters, userRealm.GetCloudAudienceURN(), &wsTrustEndpoint); err == nil {
				if samlGrant, err := wsTrustResponse.GetSAMLAssertion(&wsTrustEndpoint); err == nil {
					return req.webRequestManager.GetAccessTokenFromSamlGrant(req.authParameters, samlGrant)
				}
			}
		}
		// todo: check for ui interaction in api result...
		return nil, err
	case Managed:
		log.Println("MANAGED")
		return req.webRequestManager.GetAccessTokenFromUsernamePassword(req.authParameters)
	default:
		return nil, errors.New("Unknown account type")
	}
}
