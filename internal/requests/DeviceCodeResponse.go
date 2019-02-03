package requests

import (
	"encoding/json"
	"log"
	"time"

	"github.com/markzuber/msalgo/pkg/parameters"
)

type deviceCodeResponse struct {
	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresIn       int    `json:"expires_in"`
	Interval        int    `json:"interval"`
	Message         string `json:"message"`
}

// createDeviceCodeResponse stuff
func createDeviceCodeResponse(responseData string) (*deviceCodeResponse, error) {
	log.Println(responseData)
	dcResponse := &deviceCodeResponse{}
	var err = json.Unmarshal([]byte(responseData), dcResponse)
	if err != nil {
		return nil, err
	}

	return dcResponse, nil
}

func (dcr *deviceCodeResponse) toDeviceCodeResult(clientID string, scopes []string) *parameters.DeviceCodeResult {
	expiresOn := time.Now().UTC().Add(time.Duration(dcr.ExpiresIn) * time.Second)
	return parameters.CreateDeviceCodeResult(dcr.UserCode, dcr.DeviceCode, dcr.VerificationURL, expiresOn, dcr.Interval, dcr.Message, clientID, scopes)
}
