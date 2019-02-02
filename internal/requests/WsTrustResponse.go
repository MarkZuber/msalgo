package requests

import (
	"errors"
	"log"
)

type WsTrustResponse struct {
	responseData string
}

func CreateWsTrustResponse(responseData string) *WsTrustResponse {
	log.Println("CreateWsTrustResponse ENTERED")
	log.Println(responseData)
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

			log.Println(wsTrustResponse.responseData)
			return CreateSamlTokenInfo(SamlV2, "theassertion"), nil

			// auto tokenResponseCollection =
			//     _doc.child("s:Envelope").child("s:Body").child("trust:RequestSecurityTokenResponseCollection");
			// for (const auto& tokenResponse : tokenResponseCollection.children("trust:RequestSecurityTokenResponse"))
			// {
			//     auto token = tokenResponse.child("trust:RequestedSecurityToken");
			//     if (token.first_child() != nullptr)
			//     {
			// 		log.Println("Found valid assertion, converting to string")
			//         stringstream assertion;
			//         // We need pugixml to not add whitespace, since the assertion is signed
			//         token.child("saml:Assertion").print(assertion, "", pugi::format_raw);
			//         string samlVersion = token.child("saml:Assertion").attribute("xmlns:saml").value();
			//         if (samlVersion == "urn:oasis:names:tc:SAML:1.0:assertion")
			//         {
			//             log.Println("Retrieved WS-Trust 1.3 / SAML V1 assertion")
			//             return make_shared<SamlTokenInfo>(SamlV1, assertion.str());
			//         }
			//         if (samlVersion == "urn:oasis:names:tc:SAML:2.0:assertion")
			//         {
			//             log.Println("Retrieved WS-Trust 1.3 / SAML V2 assertion")
			//             return make_shared<SamlTokenInfo>(SamlV2, assertion.str());
			// 		}
			// 		return errors.New("Couldn't parse SAML assertion, version unknown: '%s'", samlVersion)
			//     }
			// }
		}
	default:
		return nil, errors.New("Unknown WS-Trust version")
	}
}
