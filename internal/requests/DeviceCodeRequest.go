package requests

import (
	"errors"
	"log"
	"time"

	"github.com/markzuber/msalgo/internal/msalbase"
	"github.com/markzuber/msalgo/pkg/parameters"
)

// DeviceCodeRequest stuff
type DeviceCodeRequest struct {
	webRequestManager IWebRequestManager
	authParameters    *msalbase.AuthParametersInternal
}

// CreateDeviceCodeRequest stuff
func CreateDeviceCodeRequest(
	webRequestManager IWebRequestManager,
	authParameters *msalbase.AuthParametersInternal) *DeviceCodeRequest {
	req := &DeviceCodeRequest{webRequestManager, authParameters}
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
	log.Printf("%v", deviceCodeResult)

	return req.waitForTokenResponse(deviceCodeResult)
}

func (req *DeviceCodeRequest) waitForTokenResponse(deviceCodeResult *parameters.DeviceCodeResult) (*msalbase.TokenResponse, error) {

	timeRemaining := deviceCodeResult.GetExpiresOn().Sub(time.Now().UTC())

	for timeRemaining.Seconds() > 0.0 {
		// todo: how to check for cancellation requested...

		tokenResponse, err := req.webRequestManager.GetAccessTokenFromDeviceCodeResult(req.authParameters, deviceCodeResult)
		if err != nil {
			if isErrorAuthorizationPending(err) {
				timeRemaining = deviceCodeResult.GetExpiresOn().Sub(time.Now().UTC())
			} else {
				return nil, err
			}
		} else {
			return tokenResponse, nil
		}
	}

	return nil, errors.New("Verification code expired before contacting the server")
}
