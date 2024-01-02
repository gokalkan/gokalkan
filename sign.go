package gokalkan

import (
	"encoding/base64"
	"encoding/xml"
	"strings"

	"github.com/gokalkan/gokalkan/ckalkan"
	"github.com/gokalkan/gokalkan/types"
)

// Sign подписывает данные и возвращает CMS с подписью.
func (cli *Client) Sign(input *types.SignInput) (signature []byte, err error) {
	dataB64 := base64.StdEncoding.EncodeToString(input.DataBytes)
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

	if input.WithTSP {
		flags |= ckalkan.FlagWithTimestamp
	}

	if input.IsDetached {
		flags |= ckalkan.FlagDetachedData
	}

	signatureB64, err := cli.kc.SignData("", dataB64, "", flags)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(signatureB64)
}

// SignXML подписывает данные в формате XML.
func (cli *Client) SignXML(input *types.SignXMLInput) (string, error) {
	var flags ckalkan.Flag

	if input.WithTSP {
		flags = ckalkan.FlagWithTimestamp
	}

	return cli.kc.SignXML(input.Data, "", flags, "", "", "")
}

func (cli *Client) SignWSSE(xmlData, id string) (string, error) {
	soapEnvelope := WrapWithWSSESoapEnvelope(xmlData, id)
	return cli.kc.SignWSSE(soapEnvelope, "", 0, id)
}

// SignHash подписывает hash и возвращает CMS с подписью.
func (cli *Client) SignHash(input *types.SignHashInput) (signedHash []byte, err error) {
	dataB64 := base64.StdEncoding.EncodeToString(input.InHash)
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

	if input.WithTSP {
		flags |= ckalkan.FlagWithTimestamp
	}

	if input.IsDetached {
		flags |= ckalkan.FlagDetachedData
	}

	signatureB64, err := cli.kc.SignHash(input.Algo, dataB64, flags)
	if err != nil {
		return nil, err
	}

	return base64.StdEncoding.DecodeString(signatureB64)
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
