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

	log.Println("acquiring token by username password")
	params := parameters.CreateAcquireTokenUsernamePasswordParameters("user.read", "mzuber@microsoft.com", "abc123")
	result, err := pca.AcquireTokenByUsernamePassword(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result.GetAccessToken())
}
