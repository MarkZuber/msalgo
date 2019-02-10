package msalbase

import "github.com/markzuber/msalgo/pkg/contracts"

type Account struct {
}

func CreateAccount(homeAccountID string,
	environment string,
	realm string,
	localAccountID string,
	authorityType AuthorityType,
	preferredUsername string,
	givenName string,
	familyName string,
	middleName string,
	name string,
	alternativeID string,
	rawClientInfo string,
	additionalFieldsJson string) contracts.IAccount {
	a := &Account{}
	return a
}
