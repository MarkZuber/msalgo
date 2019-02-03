package parameters

// AcquireTokenCommonParameters stuff
type AcquireTokenCommonParameters struct {
	scopes []string
}

func createAcquireTokenCommonParameters(scopes []string) *AcquireTokenCommonParameters {
	p := &AcquireTokenCommonParameters{
		scopes: scopes,
	}
	return p
}
