package gokalkan

import (
	"encoding/base64"
	"encoding/xml"
	"strings"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// Sign подписывает данные и возвращает CMS с подписью.
func (cli *Client) Sign(data []byte, isDetached, withTSP bool) (signature []byte, err error) {
	dataB64 := base64.StdEncoding.EncodeToString(data)
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

	if withTSP {
		flags |= ckalkan.FlagWithTimestamp
	}

	if isDetached {
		flags |= ckalkan.FlagDetachedData
	}

	signatureB64, err := cli.kc.SignData("", dataB64, "", flags)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(signatureB64)
}

func (cli *Client) SignDetached(data, signature []byte, withTSP bool) ([]byte, error) {
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagIn2Base64 | ckalkan.FlagOutBase64 | ckalkan.FlagDetachedData

	if withTSP {
		flags |= ckalkan.FlagWithTimestamp
	}

	signatureB64, err := cli.kc.SignData(
		base64.StdEncoding.EncodeToString(signature),
		base64.StdEncoding.EncodeToString(data),
		"",
		flags,
	)

	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(signatureB64)
}

// SignXML подписывает данные в формате XML.
func (cli *Client) SignXML(xmlData string) (string, error) {
	return cli.kc.SignXML(xmlData, "", 0, "", "", "")
}

func (cli *Client) SignWSSE(xmlData, id string) (string, error) {
	soapEnvelope := WrapWithWSSESoapEnvelope(xmlData, id)
	return cli.kc.SignWSSE(soapEnvelope, "", 0, id)
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
