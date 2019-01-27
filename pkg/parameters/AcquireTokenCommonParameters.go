package parameters

// AcquireTokenCommonParameters stuff
type AcquireTokenCommonParameters struct {
	scopes string // todo: obviously need this to be a ScopeSet custom struct or at least a collection
}

func createAcquireTokenCommonParameters(scopes string) *AcquireTokenCommonParameters {
	p := &AcquireTokenCommonParameters{
		scopes: scopes,
	}
	return p
}
