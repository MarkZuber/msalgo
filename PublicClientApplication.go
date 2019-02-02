package msalgo

import (
	"github.com/markzuber/msalgo/internal/requests"
	"github.com/markzuber/msalgo/pkg/parameters"
	"github.com/markzuber/msalgo/pkg/results"
)

// PublicClientApplication stuff
type PublicClientApplication struct {
	commonParameters  *parameters.ApplicationCommonParameters
	pcaParameters     *parameters.PublicClientApplicationParameters
	webRequestManager requests.IWebRequestManager
}

func createPublicClientApplication(builder *PublicClientApplicationBuilder) *PublicClientApplication {

	httpManager := requests.CreateHTTPManager()
	webRequestManager := requests.CreateWebRequestManager(httpManager)

	pca := &PublicClientApplication{
		commonParameters:  builder.commonParameters,
		pcaParameters:     builder.pcaParameters,
		webRequestManager: webRequestManager,
	}
	return pca
}

// AcquireTokenByUsernamePassword stuff
func (pca *PublicClientApplication) AcquireTokenByUsernamePassword(
	usernamePasswordParameters *parameters.AcquireTokenUsernamePasswordParameters) (*results.AuthenticationResult, error) {
	authParams := createAuthParametersInternal(pca.commonParameters, usernamePasswordParameters.GetCommonParameters())
	pca.pcaParameters.AugmentAuthParametersInternal(authParams)
	authParams.SetAuthorizationType(requests.UsernamePassword)
	usernamePasswordParameters.AugmentAuthParametersInternal(authParams)
	req := requests.CreateUsernamePasswordRequest(pca.webRequestManager, authParams)
	tokenResponse, err := req.Execute()
	if err == nil {
		return results.CreateAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}

func createAuthParametersInternal(
	applicationCommonParameters *parameters.ApplicationCommonParameters,
	commonParameters *parameters.AcquireTokenCommonParameters) *requests.AuthParametersInternal {
	authParams := requests.CreateAuthParametersInternal(applicationCommonParameters.GetClientID())
	return authParams
}
