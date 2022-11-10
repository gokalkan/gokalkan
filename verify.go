package gokalkan

import (
	"encoding/base64"

	"github.com/doodocs/doodocs/pkg/gokalkan/ckalkan"
)

// Verify обеспечивает проверку подписи CMS в base64.
func (cli *Client) Verify(signature []byte) (string, error) {
	signatureB64 := base64.StdEncoding.EncodeToString(signature)
	flags := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

	resp, err := cli.kc.VerifyData(signatureB64, "", "", flags)
	if err != nil {
		return "", err
	}
	return string(resp.Info), err
}

// VerifyXML обеспечивает проверку подписи данных в формате XML.
func (cli *Client) VerifyXML(signedXML string) (result string, err error) {
	return cli.kc.VerifyXML(signedXML, "", 0)
}

// VerifyDetached обеспечивает проверку отделенной подписи
// CMS (detached signature) в base64.
func (cli *Client) VerifyDetached(signature, data []byte) (string, error) {
	signatureB64 := base64.StdEncoding.EncodeToString(signature)
	dataB64 := base64.StdEncoding.EncodeToString(data)

	flags := ckalkan.FlagSignCMS
	flags |= ckalkan.FlagInBase64
	flags |= ckalkan.FlagIn2Base64
	flags |= ckalkan.FlagDetachedData

	resp, err := cli.kc.VerifyData(signatureB64, dataB64, "", flags)
	if err != nil {
		return "", err
	}
	return string(resp.Info), err
}
