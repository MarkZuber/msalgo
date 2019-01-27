package parameters

import (
	"github.com/markzuber/msalgo/internal/requests"
)

// AcquireTokenUsernamePasswordParameters stuff
type AcquireTokenUsernamePasswordParameters struct {
	commonParameters *AcquireTokenCommonParameters
	username         string
	password         string
}

// CreateAcquireTokenUsernamePasswordParameters stuff
func CreateAcquireTokenUsernamePasswordParameters(scopes string) *AcquireTokenUsernamePasswordParameters {
	p := &AcquireTokenUsernamePasswordParameters{
		commonParameters: createAcquireTokenCommonParameters(scopes),
		username:         "",
		password:         "",
	}
	return p
}

// SetUsername stuff
func (p *AcquireTokenUsernamePasswordParameters) SetUsername(username string) {
	p.username = username
}

func (p *AcquireTokenUsernamePasswordParameters) GetUsername() string {
	return p.username
}

// SetPassword stuff
func (p *AcquireTokenUsernamePasswordParameters) SetPassword(password string) {
	p.password = password
}

func (p *AcquireTokenUsernamePasswordParameters) GetPassword() string {
	return p.password
}

func (p *AcquireTokenUsernamePasswordParameters) GetCommonParameters() *AcquireTokenCommonParameters {
	return p.commonParameters
}

func (p *AcquireTokenUsernamePasswordParameters) AugmentAuthParametersInternal(authParams *requests.AuthParametersInternal) {
	authParams.SetUsername(p.GetUsername())
	authParams.SetPassword(p.GetPassword())
}
