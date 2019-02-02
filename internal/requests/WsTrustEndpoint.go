package requests

import "log"

type WsTrustEndpointVersion int

const (
	Trust2005 WsTrustEndpointVersion = iota
	Trust13
)

type WsTrustEndpoint struct {
	endpointVersion WsTrustEndpointVersion
	url             string
}

func CreateWsTrustEndpoint(endpointVersion WsTrustEndpointVersion, url string) WsTrustEndpoint {
	return WsTrustEndpoint{endpointVersion, url}
}

func (wte *WsTrustEndpoint) GetVersion() WsTrustEndpointVersion {
	return wte.endpointVersion
}

func (wte *WsTrustEndpoint) buildTokenRequestMessage(authType AuthorizationType, cloudAudienceURN string, username string, password string) string {
	var soapAction string
	var trustNamespace string
	var keyType string
	var requestType string

	if wte.endpointVersion == Trust2005 {
		log.Println("Building WS-Trust token request for v2005")
		soapAction = Trust2005Spec
		trustNamespace = "http://schemas.xmlsoap.org/ws/2005/02/trust"
		keyType = "http://schemas.xmlsoap.org/ws/2005/05/identity/NoProofKey"
		requestType = "http://schemas.xmlsoap.org/ws/2005/02/trust/Issue"
	} else {
		log.Println("Building WS-Trust token request for v1.3")
		soapAction = Trust13Spec
		trustNamespace = "http://docs.oasis-open.org/ws-sx/ws-trust/200512"
		keyType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Bearer"
		requestType = "http://docs.oasis-open.org/ws-sx/ws-trust/200512/Issue"
	}

	return soapAction + trustNamespace + keyType + requestType

	// pugi::xml_document doc;
	// {
	//     pugi::xml_node envelope = doc.append_child("s:Envelope");
	//     envelope.append_attribute("xmlns:s") = "http://www.w3.org/2003/05/soap-envelope";
	//     envelope.append_attribute("xmlns:wsa") = "http://www.w3.org/2005/08/addressing";
	//     envelope.append_attribute("xmlns:wsu") =
	//         "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd";
	//     {
	//         pugi::xml_node header = envelope.append_child("s:Header");
	//         {
	//             pugi::xml_node action = header.append_child("wsa:Action");
	//             action.append_attribute("s:mustUnderstand") = 1;
	//             action.text().set(soapAction);
	//         }
	//         {
	//             header.append_child("wsa:messageID").text() = ("urn:uuid:" + xg::newGuid().str()).c_str();
	//         }
	//         {
	//             pugi::xml_node replyTo = header.append_child("wsa:ReplyTo");
	//             {
	//                 replyTo.append_child("wsa:Address").text() = "http://www.w3.org/2005/08/addressing/anonymous";
	//             }
	//         }
	//         {
	//             pugi::xml_node to = header.append_child("wsa:To");
	//             to.append_attribute("s:mustUnderstand") = 1;
	//             to.text().set(_url.c_str());
	//         }
	//         if (authType == AuthorizationType::UsernamePassword)
	//         {
	//             AppendSecurityHeader(header, username, password);
	//         }
	//     }
	//     {
	//         pugi::xml_node body = envelope.append_child("s:Body");
	//         {
	//             pugi::xml_node requestSecurityToken = body.append_child("wst:RequestSecurityToken");
	//             requestSecurityToken.append_attribute("xmlns:wst") = trustNamespace;
	//             {
	//                 pugi::xml_node appliesTo = requestSecurityToken.append_child("wsp:AppliesTo");
	//                 appliesTo.append_attribute("xmlns:wsp") = "http://schemas.xmlsoap.org/ws/2004/09/policy";
	//                 {
	//                     pugi::xml_node endpointReference = appliesTo.append_child("wsa:EndpointReference");
	//                     {
	//                         endpointReference.append_child("wsa:Address").text() = cloudAudienceUrn.c_str();
	//                     }
	//                 }
	//             }
	//             {
	//                 requestSecurityToken.append_child("wst:KeyType").text() = keyType;
	//             }
	//             {
	//                 requestSecurityToken.append_child("wst:RequestType").text() = requestType;
	//             }
	//         }
	//     }
	// }

	// stringstream docStream;
	// doc.save(docStream, "  ", pugi::format_default | pugi::format_no_declaration);
	// return docStream.str();
}

func (wte *WsTrustEndpoint) BuildTokenRequestMessageWIA(cloudAudienceURN string) string {
	return wte.buildTokenRequestMessage(WindowsIntegratedAuth, cloudAudienceURN, "", "")
}

func (wte *WsTrustEndpoint) BuildTokenRequestMessageUsernamePassword(cloudAudienceURN string, username string, password string) string {
	return wte.buildTokenRequestMessage(UsernamePassword, cloudAudienceURN, username, password)
}

func (wte *WsTrustEndpoint) GetURL() string {
	return wte.url
}
