package main

import (
	"log"

	"github.com/markzuber/msalgo"
)

func main() {

	log.Println("creating pca")

	pcaParameters := msalgo.CreatePublicClientApplicationParameters("0615b6ca-88d4-4884-8729-b178178f7c27")
	pcaParameters.SetAadAuthority("https://login.microsoftonline.com/organizations")
	// pcaParameters.SetHttpClient()

	pca, err := msalgo.CreatePublicClientApplication(pcaParameters)
	if err != nil {
		log.Fatal(err)
	}

	{
		log.Println("acquiring token by device code")
		deviceCodeParams := msalgo.CreateAcquireTokenDeviceCodeParameters([]string{"user.read"})
		result, err := pca.AcquireTokenByDeviceCode(deviceCodeParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result.GetAccessToken())
	}

	{
		log.Println("acquiring token by username password")
		userNameParams := msalgo.CreateAcquireTokenUsernamePasswordParameters([]string{"user.read"}, "mzuber@microsoft.com", "abc123")
		result, err := pca.AcquireTokenByUsernamePassword(userNameParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result.GetAccessToken())
	}
}
