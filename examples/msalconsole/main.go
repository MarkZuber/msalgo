package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/markzuber/msalgo"
)

func createParams() *msalgo.PublicClientApplicationParameters {
	pcaParameters := msalgo.CreatePublicClientApplicationParameters("0615b6ca-88d4-4884-8729-b178178f7c27")
	pcaParameters.SetAadAuthority("https://login.microsoftonline.com/organizations")
	// pcaParameters.SetHttpClient()
	return pcaParameters
}

func acquireByDeviceCode() {
	pca, err := msalgo.CreatePublicClientApplication(createParams())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("acquiring token by device code")
	deviceCodeParams := msalgo.CreateAcquireTokenDeviceCodeParameters([]string{"user.read"})
	result, err := pca.AcquireTokenByDeviceCode(deviceCodeParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ACCESS TOKEN: " + result.GetAccessToken())
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)

	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\r\n")
	return value
}

func acquireByUsernamePassword() {

	pca, err := msalgo.CreatePublicClientApplication(createParams())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("acquiring token by username password")

	fmt.Println("Enter username: ")
	userName := readInput()
	fmt.Println("Enter password: ")
	password := readInput()

	userNameParams := msalgo.CreateAcquireTokenUsernamePasswordParameters([]string{"user.read"}, userName, password)
	result, err := pca.AcquireTokenByUsernamePassword(userNameParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ACCESS TOKEN: " + result.GetAccessToken())
}

func main() {

	log.Println("creating pca")

	// acquireByDeviceCode()
	acquireByUsernamePassword()
}
