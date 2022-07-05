package gokalkan

import (
	"encoding/xml"
	"strings"
)

func (cli *Client) SignWSSE(dataXML, id string) (signedXML string, err error) {
	return cli.kc.KCSignWSSE(WrapWithWSSESoapEnvelope(dataXML, id), "", 0, id)
}

const (
	xmlnsSOAP = "http://schemas.xmlsoap.org/soap/envelope/"
	xmlnsWSU  = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"

	replaceKey = "replace-this"
)

// soapEnvelope представляет soap:Envelope
type soapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	SOAP    string   `xml:"xmlns:soap,attr"`
	WSU     string   `xml:"xmlns:wsu,attr"`
	Body    soapBody `xml:"soap:Body"`
}

// soapBody представляет soap:Body
type soapBody struct {
	ID      string `xml:"wsu:Id,attr"`
	Content string `xml:",chardata"`
}

// WrapWithWSSESoapEnvelope оборачивает XML документ в SOAP формат, а точнее записывает
// содержимое под тегом soap:Body
func WrapWithWSSESoapEnvelope(dataXML, id string) (result string) {
	envelope := soapEnvelope{
		SOAP: xmlnsSOAP,
		WSU:  xmlnsWSU,
		Body: soapBody{
			ID:      id,
			Content: replaceKey,
		},
	}

	b, err := xml.Marshal(envelope)
	if err != nil {
		return result
	}

	result = strings.Replace(string(b), replaceKey, dataXML, 1)

	return result
}
