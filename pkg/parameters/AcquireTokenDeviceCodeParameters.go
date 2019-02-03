package parameters

import "github.com/markzuber/msalgo/internal/msalbase"

// AcquireTokenDeviceCodeParameters stuff
type AcquireTokenDeviceCodeParameters struct {
	commonParameters *AcquireTokenCommonParameters
	username         string
	password         string
}

// CreateAcquireTokenDeviceCodeParameters stuff
func CreateAcquireTokenDeviceCodeParameters(scopes []string) *AcquireTokenDeviceCodeParameters {
	p := &AcquireTokenDeviceCodeParameters{
		commonParameters: createAcquireTokenCommonParameters(scopes),
	}
	return p
}

func (p *AcquireTokenDeviceCodeParameters) GetCommonParameters() *AcquireTokenCommonParameters {
	return p.commonParameters
}

func (p *AcquireTokenDeviceCodeParameters) AugmentAuthParametersInternal(authParams *msalbase.AuthParametersInternal) {
}
