package main

import (
	"log"

	"github.com/markzuber/msalgo"
	"github.com/markzuber/msalgo/pkg/parameters"
)

func main() {

	log.Println("creating pca")
	pca, err := msalgo.CreatePublicClientApplicationBuilder("the-client-id").Build()
	if err != nil {
		log.Fatal(err)
	}

	{
		log.Println("acquiring token by device code")
		deviceCodeParams := parameters.CreateAcquireTokenDeviceCodeParameters([]string{"user.read"})
		result, err := pca.AcquireTokenByDeviceCode(deviceCodeParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result.GetAccessToken())
	}

	{
		log.Println("acquiring token by username password")
		userNameParams := parameters.CreateAcquireTokenUsernamePasswordParameters([]string{"user.read"}, "mzuber@microsoft.com", "abc123")
		result, err := pca.AcquireTokenByUsernamePassword(userNameParams)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(result.GetAccessToken())
	}
}
