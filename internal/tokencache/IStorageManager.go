package tokencache

import "github.com/markzuber/msalgo/pkg/contracts"

type CredentialType int

const (
	CredentialTypeOauth2RefreshToken CredentialType = iota
	CredentialTypeOidcIdToken
	CredentialTypeOther
)

type OperationStatusType int

const (
	OperationStatusTypeSuccess OperationStatusType = iota
	OperationStatusTypeFailure
	OperationStatusTypeRetriableError
)

type Credential struct {
}

type OperationStatus struct {
	StatusType        OperationStatusType
	Code              int
	StatusDescription string
	PlatformCode      int
	PlatformDomain    string
}

func CreateSuccessOperationStatus() *OperationStatus {
	status := &OperationStatus{StatusType: OperationStatusTypeSuccess}
	return status
}

type AppMetadata struct {
	Environment string
	ClientID    string
	FamilyID    string
}

type ReadCredentialsResponse struct {
	Credentials     []Credential
	OperationStatus *OperationStatus
}

type ReadAccountsResponse struct {
	Accounts        []contracts.IAccount
	OperationStatus *OperationStatus
}

type ReadAccountResponse struct {
	Account         contracts.IAccount
	OperationStatus *OperationStatus
}

type IStorageManager interface {
	ReadCredentials(
		correlationID string,
		homeAccountID string,
		environment string,
		realm string,
		clientID string,
		familyID string,
		target string,
		types map[CredentialType]bool) (*ReadCredentialsResponse, error)

	WriteCredentials(correlationID string, credentials []Credential) (*OperationStatus, error)

	DeleteCredentials(
		correlationId string,
		homeAccountId string,
		environment string,
		realm string,
		clientID string,
		familyID string,
		target string,
		types map[CredentialType]bool) (*OperationStatus, error)

	ReadAllAccounts(correlationID string) (*ReadAccountsResponse, error)

	ReadAccount(
		correlationID string,
		homeAccountID string,
		environment string,
		realm string) (*ReadAccountResponse, error)

	WriteAccount(correlationID string, account contracts.IAccount) (*OperationStatus, error)

	DeleteAccount(
		correlationID string,
		homeAccountID string,
		environment string,
		realm string) (*OperationStatus, error)

	DeleteAccounts(correlationID string, homeAccountID string, environment string) (*OperationStatus, error)
	ReadAppMetadata(environment string, clientID string) (*AppMetadata, error)
	WriteAppMetadata(appMetadata *AppMetadata) error
}
