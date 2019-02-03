package msalgo

import (
	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/internal/requests"
	"github.com/markzuber/msalgo/pkg/parameters"
)

// PublicClientApplication is used to acquire tokens in desktop or mobile applications (Desktop / UWP / Xamarin.iOS / Xamarin.Android).
// public client applications are not trusted to safely keep application secrets, and therefore they only access Web APIs in the name of the user only
// (they only support public client flows). For details see https://aka.ms/msal-net-client-applications
type PublicClientApplication struct {
	commonParameters  *parameters.ApplicationCommonParameters
	pcaParameters     *parameters.PublicClientApplicationParameters
	webRequestManager requests.IWebRequestManager
}

func createPublicClientApplication(builder *PublicClientApplicationBuilder) *PublicClientApplication {

	httpManager := msalbase.CreateHTTPManager()
	webRequestManager := requests.CreateWebRequestManager(httpManager)

	pca := &PublicClientApplication{
		commonParameters:  builder.commonParameters,
		pcaParameters:     builder.pcaParameters,
		webRequestManager: webRequestManager,
	}
	return pca
}

// AcquireTokenByUsernamePassword is a non-interactive request to acquire a security token from the authority, via Username/Password Authentication.
func (pca *PublicClientApplication) AcquireTokenByUsernamePassword(
	usernamePasswordParameters *parameters.AcquireTokenUsernamePasswordParameters) (*AuthenticationResult, error) {
	authParams := createAuthParametersInternal(pca.commonParameters, usernamePasswordParameters.GetCommonParameters())
	pca.pcaParameters.AugmentAuthParametersInternal(authParams)
	authParams.SetAuthorizationType(msalbase.UsernamePassword)
	usernamePasswordParameters.AugmentAuthParametersInternal(authParams)
	req := requests.CreateUsernamePasswordRequest(pca.webRequestManager, authParams)
	tokenResponse, err := req.Execute()
	if err == nil {
		return createAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}

func (pca *PublicClientApplication) AcquireTokenByDeviceCode(
	deviceCodeParameters *parameters.AcquireTokenDeviceCodeParameters) (*AuthenticationResult, error) {
	authParams := createAuthParametersInternal(pca.commonParameters, deviceCodeParameters.GetCommonParameters())
	pca.pcaParameters.AugmentAuthParametersInternal(authParams)
	authParams.SetAuthorizationType(msalbase.DeviceCode)
	deviceCodeParameters.AugmentAuthParametersInternal(authParams)
	req := requests.CreateDeviceCodeRequest(pca.webRequestManager, authParams)
	tokenResponse, err := req.Execute()
	if err == nil {
		return createAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}

func createAuthParametersInternal(
	applicationCommonParameters *parameters.ApplicationCommonParameters,
	commonParameters *parameters.AcquireTokenCommonParameters) *msalbase.AuthParametersInternal {
	authParams := msalbase.CreateAuthParametersInternal(applicationCommonParameters.GetClientID())
	return authParams
}
