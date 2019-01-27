package msalgo

import (
	"github.com/markzuber/msalgo/internal/requests"
	"github.com/markzuber/msalgo/pkg/parameters"
	"github.com/markzuber/msalgo/pkg/results"
)

// PublicClientApplication stuff
type PublicClientApplication struct {
	commonParameters *parameters.ApplicationCommonParameters
	pcaParameters    *parameters.PublicClientApplicationParameters
}

func createPublicClientApplication(
	builder *PublicClientApplicationBuilder) *PublicClientApplication {
	pca := &PublicClientApplication{
		commonParameters: builder.commonParameters,
		pcaParameters:    builder.pcaParameters,
	}
	return pca
}

// AcquireTokenInteractive stuff
func (pca *PublicClientApplication) AcquireTokenInteractive(
	interactiveParameters *parameters.AcquireTokenInteractiveParameters) (*results.AuthenticationResult, error) {
	req := requests.CreateInteractiveRequest(interactiveParameters)
	return req.Execute()
}
