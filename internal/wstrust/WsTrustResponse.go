package wstrust

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
)

type WsTrustResponse struct {
	responseData string
}

func CreateWsTrustResponse(responseData string) *WsTrustResponse {
	log.Println("CreateWsTrustResponse ENTERED")
	// log.Println(responseData)
	response := &WsTrustResponse{responseData}
	return response

	// pugi::xml_parse_result result = _doc.load_string(response.c_str());
	// if (!result)
	// {
	//     MSAL_THROW(UNTAGGED, Status::Unexpected, result.status, "Failed to parse SAML response '%s', response");
	// }

	// auto fault = _doc.child("s:Envelope").child("s:Body").child("s:Fault");
	// if (fault != nullptr)
	// {
	//     MSAL_THROW(
	//         UNTAGGED,
	//         Status::Unexpected,
	//         "SAML assertion indicates error: Code '%s' Subcode '%s' Reason '%s'",
	//         fault.child("s:Code").child_value("s:Value"),
	//         fault.child("s:Code").child("s:Subcode").child_value("s:Value"),
	//         fault.child("s:Reason").child_value("s:Text"));
	// }
}

func (wsTrustResponse *WsTrustResponse) GetSAMLAssertion(endpoint *WsTrustEndpoint) (*SamlTokenInfo, error) {
	switch endpoint.GetVersion() {
	case Trust2005:
		return nil, errors.New("WS Trust 2005 support is not implemented")
	case Trust13:
		{
			log.Println("Extracting assertion from WS-Trust 1.3 token:")

			samldefinitions := &samldefinitions{}
			var err = xml.Unmarshal([]byte(wsTrustResponse.responseData), samldefinitions)
			if err != nil {
				return nil, err
			}

			for _, tokenResponse := range samldefinitions.Body.RequestSecurityTokenResponseCollection.RequestSecurityTokenResponse {
				token := tokenResponse.RequestedSecurityToken
				if token.Assertion.XMLName.Local != "" {
					// log.Println("Found valid assertion")
					assertion := token.AssertionRawXML

					samlVersion := token.Assertion.Saml
					if samlVersion == "urn:oasis:names:tc:SAML:1.0:assertion" {
						log.Println("Retrieved WS-Trust 1.3 / SAML V1 assertion")
						return CreateSamlTokenInfo(SamlV1, assertion), nil
					}
					if samlVersion == "urn:oasis:names:tc:SAML:2.0:assertion" {
						log.Println("Retrieved WS-Trust 1.3 / SAML V2 assertion")
						return CreateSamlTokenInfo(SamlV2, assertion), nil
					}

					return nil, fmt.Errorf("Couldn't parse SAML assertion, version unknown: '%s'", samlVersion)
				}
			}

			return nil, errors.New("Couldn't find SAML assertion")
		}
	default:
		return nil, errors.New("Unknown WS-Trust version")
	}
}
