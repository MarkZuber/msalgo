package requests

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/markzuber/msalgo/pkg/parameters"
)

type deviceCodeResponse struct {
	Error            string `json:"error"`
	SubError         string `json:"suberror"`
	ErrorDescription string `json:"error_description"`
	ErrorCodes       []int  `json:"error_codes"`
	CorrelationID    string `json:"correlation_id"`
	Claims           string `json:"claims"`

	UserCode        string `json:"user_code"`
	DeviceCode      string `json:"device_code"`
	VerificationURL string `json:"verification_url"`
	ExpiresInStr    string `json:"expires_in"`
	IntervalStr     string `json:"interval"`
	Message         string `json:"message"`

	ExpiresIn int
	Interval  int
}

// createDeviceCodeResponse stuff
func createDeviceCodeResponse(responseData string) (*deviceCodeResponse, error) {
	log.Println(responseData)
	dcResponse := &deviceCodeResponse{}
	var err = json.Unmarshal([]byte(responseData), dcResponse)
	if err != nil {
		return nil, err
	}

	expiresIn, err := strconv.Atoi(dcResponse.ExpiresInStr)
	if err != nil {
		return nil, err
	}
	dcResponse.ExpiresIn = expiresIn

	interval, err := strconv.Atoi(dcResponse.IntervalStr)
	if err != nil {
		return nil, err
	}
	dcResponse.Interval = interval

	return dcResponse, nil
}

func (dcr *deviceCodeResponse) toDeviceCodeResult(clientID string, scopes []string) *parameters.DeviceCodeResult {
	expiresOn := time.Now().UTC().Add(time.Duration(dcr.ExpiresIn) * time.Second)
	return parameters.CreateDeviceCodeResult(dcr.UserCode, dcr.DeviceCode, dcr.VerificationURL, expiresOn, dcr.Interval, dcr.Message, clientID, scopes)
}
