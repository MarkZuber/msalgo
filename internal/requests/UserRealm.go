package requests

import "encoding/json"

type UserRealmAccountType int

const (
	Unknown UserRealmAccountType = iota
	Federated
	Managed
)

type UserRealm struct {
	accountType       string `json:"account_type"`
	domainName        string `json:"domain_name"`
	cloudInstanceNmae string `json:"cloud_instance_name"`
	cloudAudienceURN  string `json:"cloud_audience_urn"`

	// required if accountType is Federated
	federationProtocol    string `json:"federation_protocol"`
	federationMetadataURL string `json:"federation_metadata_url"`
}

// CreateUserRealm stuff
func CreateUserRealm(responseData string) (*UserRealm, error) {
	var userRealm *UserRealm
	var err = json.Unmarshal([]byte(responseData), userRealm)
	if err != nil {
		return nil, err
	}

	if userRealm.GetAccountType() == Federated {
		// assert federationProtocol and federationMetadataURL are set/valid/non-null
	}

	// assert domainName, cloudInstanceName, cloudInstanceUrn are set/valid/non-null

	return userRealm, nil
}

func (u *UserRealm) GetAccountType() UserRealmAccountType {
	if u.accountType == "Federated" {
		return Federated
	}
	if u.accountType == "Managed" {
		return Managed
	}
	return Unknown
}

func (u *UserRealm) GetFederationMetadataURL() string {
	return u.federationMetadataURL
}

func (u *UserRealm) GetCloudAudienceURN() string {
	return u.cloudAudienceURN
}
