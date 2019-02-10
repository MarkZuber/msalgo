package requests

import (
	"errors"
	"log"

	"github.com/markzuber/msalgo/internal/tokencache"

	"github.com/markzuber/msalgo/internal/msalbase"
)

// UsernamePasswordRequest stuff
type UsernamePasswordRequest struct {
	webRequestManager IWebRequestManager
	cacheManager      tokencache.ICacheManager
	authParameters    *msalbase.AuthParametersInternal
}

// CreateUsernamePasswordRequest stuff
func CreateUsernamePasswordRequest(
	webRequestManager IWebRequestManager,
	cacheManager tokencache.ICacheManager,
	authParameters *msalbase.AuthParametersInternal) *UsernamePasswordRequest {
	req := &UsernamePasswordRequest{webRequestManager, cacheManager, authParameters}
	return req
}

// Execute stuff
func (req *UsernamePasswordRequest) Execute() (*msalbase.TokenResponse, error) {

	resolutionManager := CreateAuthorityEndpointResolutionManager(req.webRequestManager)
	endpoints, err := resolutionManager.ResolveEndpoints(req.authParameters.GetAuthorityInfo(), "")
	if err != nil {
		return nil, err
	}

	req.authParameters.SetAuthorityEndpoints(endpoints)

	userRealm, err := req.webRequestManager.GetUserRealm(req.authParameters)
	if err != nil {
		return nil, err
	}

	// log.Println("got user realm")

	switch accountType := userRealm.GetAccountType(); accountType {
	case msalbase.Federated:
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
	case msalbase.Managed:
		log.Println("MANAGED")
		return req.webRequestManager.GetAccessTokenFromUsernamePassword(req.authParameters)
	default:
		return nil, errors.New("Unknown account type")
	}
}
