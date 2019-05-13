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
	cacheManager      msalbase.ICacheManager
}

// CreatePublicClientApplication stuff
func CreatePublicClientApplication(pcaParameters *PublicClientApplicationParameters) (*PublicClientApplication, error) {
	err := pcaParameters.validate()
	if err != nil {
		return nil, err
	}

	httpManager := msalbase.CreateHTTPManager()
	webRequestManager := requests.CreateWebRequestManager(httpManager)

	// todo: check parameters for whether persistent cache is desired, or self-caching (callback to byte array read/write)
	storageManager := tokencache.CreateStorageManager()
	cacheManager := tokencache.CreateCacheManager(storageManager)

	pca := &PublicClientApplication{pcaParameters, webRequestManager, cacheManager}
	return pca, nil
}

// AcquireTokenSilent stuff
func (pca *PublicClientApplication) AcquireTokenSilent(
	silentParameters *AcquireTokenSilentParameters) (IAuthenticationResult, error) {
	authParams := pca.pcaParameters.createAuthenticationParameters()
	silentParameters.augmentAuthenticationParameters(authParams)

	storageTokenResponse, err := pca.cacheManager.TryReadCache(authParams)
	if err != nil {
		return nil, err
	}

	if storageTokenResponse != nil {
		return msalbase.CreateAuthenticationResultFromStorageTokenResponse(storageTokenResponse)
	}

	req := requests.CreateRefreshTokenExchangeRequest(pca.webRequestManager, pca.cacheManager, authParams)
	return pca.executeTokenRequestWithCacheWrite(req, authParams)
}

// AcquireTokenByUsernamePassword is a non-interactive request to acquire a security token from the authority, via Username/Password Authentication.
func (pca *PublicClientApplication) AcquireTokenByUsernamePassword(
	usernamePasswordParameters *AcquireTokenUsernamePasswordParameters) (IAuthenticationResult, error) {
	authParams := pca.pcaParameters.createAuthenticationParameters()
	usernamePasswordParameters.augmentAuthenticationParameters(authParams)

	req := requests.CreateUsernamePasswordRequest(pca.webRequestManager, pca.cacheManager, authParams)
	return pca.executeTokenRequestWithoutCacheWrite(req, authParams)
}

// AcquireTokenByDeviceCode stuff
func (pca *PublicClientApplication) AcquireTokenByDeviceCode(
	deviceCodeParameters *AcquireTokenDeviceCodeParameters) (IAuthenticationResult, error) {
	authParams := pca.pcaParameters.createAuthenticationParameters()
	deviceCodeParameters.augmentAuthenticationParameters(authParams)

	req := requests.CreateDeviceCodeRequest(pca.webRequestManager, pca.cacheManager, authParams)
	return pca.executeTokenRequestWithoutCacheWrite(req, authParams)
}

// AcquireTokenInteractive stuff
func (pca *PublicClientApplication) AcquireTokenInteractive(
	interactiveParams *AcquireTokenInteractiveParameters) (IAuthenticationResult, error) {
	authParams := pca.pcaParameters.createAuthenticationParameters()
	interactiveParams.augmentAuthenticationParameters(authParams)

	req := requests.CreateInteractiveRequest(pca.webRequestManager, pca.cacheManager, authParams)
	return pca.executeTokenRequestWithCacheWrite(req, authParams)
}

// executeTokenRequestWithoutCacheWrite stuff
func (pca *PublicClientApplication) executeTokenRequestWithoutCacheWrite(
	req requests.ITokenRequester,
	authParams *msalbase.AuthParametersInternal) (IAuthenticationResult, error) {
	tokenResponse, err := req.Execute()
	if err == nil {
		// todo: is account being nil proper here?
		return msalbase.CreateAuthenticationResult(tokenResponse, nil), nil
	}
	return nil, err
}

// executeTokenRequestWithCacheWrite stuff
func (pca *PublicClientApplication) executeTokenRequestWithCacheWrite(
	req requests.ITokenRequester,
	authParams *msalbase.AuthParametersInternal) (IAuthenticationResult, error) {
	tokenResponse, err := req.Execute()
	if err == nil {
		account, err := pca.cacheManager.CacheTokenResponse(authParams, tokenResponse)
		if err != nil {
			return nil, err
		}
		return msalbase.CreateAuthenticationResult(tokenResponse, account), nil
	}
	return nil, err
}
