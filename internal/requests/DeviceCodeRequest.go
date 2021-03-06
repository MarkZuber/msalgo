package requests

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/markzuber/msalgo/internal/msalbase"
)

// DeviceCodeRequest stuff
type DeviceCodeRequest struct {
	webRequestManager IWebRequestManager
	cacheManager      msalbase.ICacheManager
	authParameters    *msalbase.AuthParametersInternal
}

// CreateDeviceCodeRequest stuff
func CreateDeviceCodeRequest(
	webRequestManager IWebRequestManager,
	cacheManager msalbase.ICacheManager,
	authParameters *msalbase.AuthParametersInternal) *DeviceCodeRequest {
	req := &DeviceCodeRequest{webRequestManager, cacheManager, authParameters}
	return req
}

// Execute stuff
func (req *DeviceCodeRequest) Execute() (*msalbase.TokenResponse, error) {

	// resolve authority endpoints

	deviceCodeResult, err := req.webRequestManager.GetDeviceCodeResult(req.authParameters)
	if err != nil {
		return nil, err
	}

	// fire deviceCodeResult up to user
	log.Infof("%v", deviceCodeResult)

	return req.waitForTokenResponse(deviceCodeResult)
}

func (req *DeviceCodeRequest) waitForTokenResponse(deviceCodeResult *msalbase.DeviceCodeResult) (*msalbase.TokenResponse, error) {

	timeRemaining := deviceCodeResult.GetExpiresOn().Sub(time.Now().UTC())

	for timeRemaining.Seconds() > 0.0 {
		// todo: how to check for cancellation requested...

		// todo: learn more about go error handling so that this is managed through error flow and not parsing the token response...

		tokenResponse, err := req.webRequestManager.GetAccessTokenFromDeviceCodeResult(req.authParameters, deviceCodeResult)
		if err != nil {
			if isErrorAuthorizationPending(err) {
				timeRemaining = deviceCodeResult.GetExpiresOn().Sub(time.Now().UTC())
			} else {
				return nil, err
			}
		} else {
			if tokenResponse.IsAuthorizationPending() {
				timeRemaining = deviceCodeResult.GetExpiresOn().Sub(time.Now().UTC())
			} else {
				return tokenResponse, nil
			}
		}

		time.Sleep(5 * time.Second)
	}

	return nil, errors.New("Verification code expired before contacting the server")
}
