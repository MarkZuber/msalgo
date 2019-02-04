package msalgo

import (
	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/internal/requests"
	"github.com/markzuber/msalgo/internal/tokencache"
)

// PublicClientApplication is used to acquire tokens in desktop or mobile applications (Desktop / UWP / Xamarin.iOS / Xamarin.Android).
// public client applications are not trusted to safely keep application secrets, and therefore they only access Web APIs in the name of the user only
// (they only support public client flows). For details see https://aka.ms/msal-net-client-applications
type PublicClientApplication struct {
	pcaParameters     *PublicClientApplicationParameters
	webRequestManager requests.IWebRequestManager
	storageManager    tokencache.IStorageManager
}

func CreatePublicClientApplication(pcaParameters *PublicClientApplicationParameters) (*PublicClientApplication, error) {
	err := pcaParameters.validate()
	if err != nil {
		return nil, err
	}

	httpManager := msalbase.CreateHTTPManager()
	webRequestManager := requests.CreateWebRequestManager(httpManager)
	storageManager := tokencache.CreateStorageManager()

	pca := &PublicClientApplication{pcaParameters, webRequestManager, storageManager}
	return pca, nil
}

// AcquireTokenByUsernamePassword is a non-interactive request to acquire a security token from the authority, via Username/Password Authentication.
func (pca *PublicClientApplication) AcquireTokenByUsernamePassword(usernamePasswordParameters *AcquireTokenUsernamePasswordParameters) (*AuthenticationResult, error) {

	authParams := pca.pcaParameters.createAuthenticationParameters()
	usernamePasswordParameters.augmentAuthenticationParameters(authParams)

	cacheManager := tokencache.CreateCacheManager(pca.storageManager, authParams)

	req := requests.CreateUsernamePasswordRequest(pca.webRequestManager, cacheManager, authParams)
	tokenResponse, err := req.Execute()
	if err == nil {
		return createAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}

func (pca *PublicClientApplication) AcquireTokenByDeviceCode(deviceCodeParameters *AcquireTokenDeviceCodeParameters) (*AuthenticationResult, error) {
	authParams := pca.pcaParameters.createAuthenticationParameters()
	deviceCodeParameters.augmentAuthenticationParameters(authParams)

	cacheManager := tokencache.CreateCacheManager(pca.storageManager, authParams)

	req := requests.CreateDeviceCodeRequest(pca.webRequestManager, cacheManager, authParams)
	tokenResponse, err := req.Execute()
	if err == nil {
		return createAuthenticationResult(tokenResponse), nil
	}
	return nil, err
}
