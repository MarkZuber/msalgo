package parameters

import "github.com/markzuber/msalgo/internal/msalbase"

// PublicClientApplicationParameters stuff
type PublicClientApplicationParameters struct {
}

// CreatePublicClientApplicationParameters stuff
func CreatePublicClientApplicationParameters() *PublicClientApplicationParameters {
	p := &PublicClientApplicationParameters{}
	return p
}

func (p *PublicClientApplicationParameters) AugmentAuthParametersInternal(authParams *msalbase.AuthParametersInternal) {

}
