package requests

import (
	"encoding/xml"
	"log"
	"time"

	uuid "github.com/twinj/uuid"
)

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

type wsTrustTokenRequestEnvelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	S       string   `xml:"s,attr"`
	Wsa     string   `xml:"wsa,attr"`
	Wsu     string   `xml:"wsu,attr"`
	Header  struct {
		Text   string `xml:",chardata"`
		Action struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
		} `xml:"Action"`
		MessageID struct {
			Text string `xml:",chardata"`
		} `xml:"messageID"`
		ReplyTo struct {
			Text    string `xml:",chardata"`
			Address struct {
				Text string `xml:",chardata"`
			} `xml:"Address"`
		} `xml:"ReplyTo"`
		To struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
		} `xml:"To"`
		Security struct {
			Text           string `xml:",chardata"`
			MustUnderstand string `xml:"mustUnderstand,attr"`
			Wsse           string `xml:"wsse,attr"`
			Timestamp      struct {
				Text    string `xml:",chardata"`
				ID      string `xml:"Id,attr"`
				Created struct {
					Text string `xml:",chardata"`
				} `xml:"Created"`
				Expires struct {
					Text string `xml:",chardata"`
				} `xml:"Expires"`
			} `xml:"Timestamp"`
			UsernameToken struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"Id,attr"`
				Username struct {
					Text string `xml:",chardata"`
				} `xml:"Username"`
				Password struct {
					Text string `xml:",chardata"`
				} `xml:"Password"`
			} `xml:"UsernameToken"`
		} `xml:"Security"`
	} `xml:"Header"`
	Body struct {
		Text                 string `xml:",chardata"`
		RequestSecurityToken struct {
			Text      string `xml:",chardata"`
			Wst       string `xml:"wst,attr"`
			AppliesTo struct {
				Text              string `xml:",chardata"`
				Wsp               string `xml:"wsp,attr"`
				EndpointReference struct {
					Text    string `xml:",chardata"`
					Address struct {
						Text string `xml:",chardata"`
					} `xml:"Address"`
				} `xml:"EndpointReference"`
			} `xml:"AppliesTo"`
			KeyType struct {
				Text string `xml:",chardata"`
			} `xml:"KeyType"`
			RequestType struct {
				Text string `xml:",chardata"`
			} `xml:"RequestType"`
		} `xml:"RequestSecurityToken"`
	} `xml:"Body"`
}

func buildTimeString(t time.Time) string {
	// Golang time formats are weird: https://stackoverflow.com/questions/20234104/how-to-format-current-time-using-a-yyyymmddhhmmss-format
	return t.Format("2006-01-02T15:04:05.000Z")
}

func (wte *WsTrustEndpoint) buildTokenRequestMessage(authType AuthorizationType, cloudAudienceURN string, username string, password string) (string, error) {
	var soapAction string
	var trustNamespace string
	var keyType string
	var requestType string

	createdTime := time.Now().UTC()
	expiresTime := createdTime.Add(10 * time.Minute)

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

	var envelope wsTrustTokenRequestEnvelope

	messageUUID := uuid.NewV4()

	envelope.S = "http://www.w3.org/2003/05/soap-envelope"
	envelope.Wsa = "http://www.w3.org/2005/08/addressing"
	envelope.Wsu = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"

	envelope.Header.Action.MustUnderstand = "1"
	envelope.Header.Action.Text = soapAction
	envelope.Header.MessageID.Text = "urn:uuid:" + messageUUID.String()
	envelope.Header.ReplyTo.Address.Text = "http://www.w3.org/2005/08/addressing/anonymous"
	envelope.Header.To.MustUnderstand = "1"
	envelope.Header.To.Text = wte.url

	// note: uuid on golang: https://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
	// using "github.com/twinj/uuid"

	if authType == UsernamePassword {

		endpointUUID := uuid.NewV4()

		var trustID string
		if wte.endpointVersion == Trust2005 {
			trustID = "UnPwSecTok2005-" + endpointUUID.String()
		} else {
			trustID = "UnPwSecTok13-" + endpointUUID.String()
		}

		envelope.Header.Security.MustUnderstand = "1"
		envelope.Header.Security.Wsse = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
		envelope.Header.Security.Timestamp.ID = "MSATimeStamp"
		envelope.Header.Security.Timestamp.Created.Text = buildTimeString(createdTime)
		envelope.Header.Security.Timestamp.Expires.Text = buildTimeString(expiresTime)
		envelope.Header.Security.UsernameToken.ID = trustID
		envelope.Header.Security.UsernameToken.Username.Text = username
		envelope.Header.Security.UsernameToken.Password.Text = password
	}

	envelope.Body.RequestSecurityToken.Wst = trustNamespace
	envelope.Body.RequestSecurityToken.AppliesTo.Wsp = "http://schemas.xmlsoap.org/ws/2004/09/policy"
	envelope.Body.RequestSecurityToken.AppliesTo.EndpointReference.Address.Text = cloudAudienceURN
	envelope.Body.RequestSecurityToken.KeyType.Text = keyType
	envelope.Body.RequestSecurityToken.RequestType.Text = requestType

	output, err := xml.Marshal(envelope)
	if err != nil {
		return "", err
	}

	log.Println(string(output))

	return string(output), nil

	// return "", soapAction + trustNamespace + keyType + requestType

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

func (wte *WsTrustEndpoint) BuildTokenRequestMessageWIA(cloudAudienceURN string) (string, error) {
	return wte.buildTokenRequestMessage(WindowsIntegratedAuth, cloudAudienceURN, "", "")
}

func (wte *WsTrustEndpoint) BuildTokenRequestMessageUsernamePassword(cloudAudienceURN string, username string, password string) (string, error) {
	return wte.buildTokenRequestMessage(UsernamePassword, cloudAudienceURN, username, password)
}

func (wte *WsTrustEndpoint) GetURL() string {
	return wte.url
}
