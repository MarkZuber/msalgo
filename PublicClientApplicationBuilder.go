package msalgo

import (
	"github.com/markzuber/msalgo/pkg/parameters"
)

// PublicClientApplicationBuilder stuff
type PublicClientApplicationBuilder struct {
	commonParameters *parameters.ApplicationCommonParameters
	pcaParameters    *parameters.PublicClientApplicationParameters
}

// CreatePublicClientApplicationBuilder stuff
func CreatePublicClientApplicationBuilder(clientID string) *PublicClientApplicationBuilder {
	pca := &PublicClientApplicationBuilder{
		commonParameters: parameters.CreateApplicationCommonParameters(clientID),
		pcaParameters:    parameters.CreatePublicClientApplicationParameters(),
	}

	return pca
}

// WithAadAuthority stuff
func (b *PublicClientApplicationBuilder) WithAadAuthority() *PublicClientApplicationBuilder {
	b.commonParameters.SetAuthorityURI("https://login.microsoftonline.com/organizations")
	return b
}

func (b *PublicClientApplicationBuilder) validate() error {
	return nil
}

// Build stuff
func (b *PublicClientApplicationBuilder) Build() (*PublicClientApplication, error) {
	err := b.validate()
	if err != nil {
		return nil, err
	}

	return createPublicClientApplication(b), nil
}
