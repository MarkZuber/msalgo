package requests

import (
	"github.com/markzuber/msalgo/pkg/parameters"
	"github.com/markzuber/msalgo/pkg/results"
)

// InteractiveRequest stuff
type InteractiveRequest struct {
	parameters *parameters.AcquireTokenInteractiveParameters
}

// CreateInteractiveRequest stuff
func CreateInteractiveRequest(parameters *parameters.AcquireTokenInteractiveParameters) *InteractiveRequest {
	r := &InteractiveRequest{parameters: parameters}
	return r
}

func (r *InteractiveRequest) Execute() (*results.AuthenticationResult, error) {
	return results.CreateAuthenticationResult("the access token"), nil
}
