package msalgo

import (
	"github.com/markzuber/msalgo/pkg/parameters"
)

// PublicClientApplicationBuilder is used to construct a PublicClientApplication from a variety of parameters.
type PublicClientApplicationBuilder struct {
	commonParameters *parameters.ApplicationCommonParameters
	pcaParameters    *parameters.PublicClientApplicationParameters
}

// CreatePublicClientApplicationBuilder creates a PublicClientApplicationBuilder from a clientID.
func CreatePublicClientApplicationBuilder(clientID string) *PublicClientApplicationBuilder {
	pca := &PublicClientApplicationBuilder{
		commonParameters: parameters.CreateApplicationCommonParameters(clientID),
		pcaParameters:    parameters.CreatePublicClientApplicationParameters(),
	}

	return pca
}

// WithAadAuthority Adds a known Azure AD authority to the application to sign-in users from a single
// organization (single tenant application) described by its domain name. See https://aka.ms/msal-net-application-configuration.
func (b *PublicClientApplicationBuilder) WithAadAuthority() *PublicClientApplicationBuilder {
	b.commonParameters.SetAuthorityURI("https://login.microsoftonline.com/organizations")
	return b
}

func (b *PublicClientApplicationBuilder) validate() error {
	return nil
}

// Build constructs the PublicClientApplication.
func (b *PublicClientApplicationBuilder) Build() (*PublicClientApplication, error) {
	err := b.validate()
	if err != nil {
		return nil, err
	}

	return createPublicClientApplication(b), nil
}
