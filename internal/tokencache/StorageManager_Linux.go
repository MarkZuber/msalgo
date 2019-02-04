// +build linux

package tokencache

import (
	"errors"

	"github.com/markzuber/msalgo/pkg/contracts"
)

type linuxStorageManager struct {
}

func CreateStorageManager() IStorageManager {
	mgr := &linuxStorageManager{}
	return mgr
}

func (m *linuxStorageManager) ReadCredentials(
	correlationID string,
	homeAccountID string,
	environment string,
	realm string,
	clientID string,
	familyID string,
	target string,
	types map[CredentialType]bool) (*ReadCredentialsResponse, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) WriteCredentials(correlationID string, credentials []Credential) (*OperationStatus, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) DeleteCredentials(
	correlationID string,
	homeAccountID string,
	environment string,
	realm string,
	clientID string,
	familyID string,
	target string,
	types map[CredentialType]bool) (*OperationStatus, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) ReadAllAccounts(correlationID string) (*ReadAccountsResponse, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) ReadAccount(
	correlationID string,
	homeAccountID string,
	environment string,
	realm string) (*ReadAccountResponse, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) WriteAccount(correlationID string, account contracts.IAccount) (*OperationStatus, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) DeleteAccount(
	correlationID string,
	homeAccountID string,
	environment string,
	realm string) (*OperationStatus, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) DeleteAccounts(correlationID string, homeAccountID string, environment string) (*OperationStatus, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) ReadAppMetadata(environment string, clientID string) (*AppMetadata, error) {
	return nil, errors.New("not implemented")
}

func (m *linuxStorageManager) WriteAppMetadata(appMetadata *AppMetadata) error {
	return errors.New("not implemented")
}
