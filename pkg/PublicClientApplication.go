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

func createPublicClientApplication(
	builder *PublicClientApplicationBuilder) *PublicClientApplication {
	pca := &PublicClientApplication{
		commonParameters: builder.commonParameters,
		pcaParameters:    builder.pcaParameters,
	}
	return pca
}

// AcquireTokenByUsernamePassword stuff
func (pca *PublicClientApplication) AcquireTokenByUsernamePassword(
	usernamePasswordParameters *parameters.AcquireTokenUsernamePasswordParameters) (*results.AuthenticationResult, error) {
	authParams := createAuthParametersInternal(pca.commonParameters, usernamePasswordParameters.GetCommonParameters())
	pca.pcaParameters.AugmentAuthParametersInternal(authParams)
	usernamePasswordParameters.AugmentAuthParametersInternal(authParams)
	req := requests.CreateUsernamePasswordRequest(pca.webRequestManager, authParams)
	if tokenResponse, err := req.Execute(); err == nil {
		return results.CreateAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}

func createAuthParametersInternal(
	applicationCommonParameters *parameters.ApplicationCommonParameters,
	commonParameters *parameters.AcquireTokenCommonParameters) *AuthParametersInternal {
	authParams := requests.CreateAuthParametersInternal(applicationCommonParameters.GetClientID())
	return authParams
}
