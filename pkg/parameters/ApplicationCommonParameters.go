package parameters

// ApplicationCommonParameters stuff
type ApplicationCommonParameters struct {
	clientID     string
	authorityURI string
}

// CreateApplicationCommonParameters stuff
func CreateApplicationCommonParameters(clientID string) *ApplicationCommonParameters {
	p := &ApplicationCommonParameters{
		clientID:     clientID,
		authorityURI: "",
	}
	return p
}

// SetAuthorityURI stuff
func (p *ApplicationCommonParameters) SetAuthorityURI(authorityURI string) {
	p.authorityURI = authorityURI
}

func (p *ApplicationCommonParameters) GetClientID() string {
	return p.clientID
}
