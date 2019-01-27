package parameters

// AcquireTokenInteractiveParameters stuff
type AcquireTokenInteractiveParameters struct {
	commonParameters *AcquireTokenCommonParameters
	loginHint        string
}

// CreateAcquireTokenInteractiveParameters stuff
func CreateAcquireTokenInteractiveParameters(scopes string) *AcquireTokenInteractiveParameters {
	p := &AcquireTokenInteractiveParameters{
		commonParameters: createAcquireTokenCommonParameters(scopes),
		loginHint:        "",
	}
	return p
}

// SetLoginHint stuff
func (p *AcquireTokenInteractiveParameters) SetLoginHint(loginHint string) {
	p.loginHint = loginHint
}
