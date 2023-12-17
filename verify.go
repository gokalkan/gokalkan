package gokalkan

import (
	"encoding/base64"
	"github.com/gokalkan/gokalkan/types"

	"github.com/gokalkan/gokalkan/ckalkan"
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
func (cli *Client) VerifyXML(input *types.VerifyInput) (result string, err error) {
	var (
		flags     ckalkan.Flag
		signedXML string
	)

	signedXML = string(input.SignatureBytes)

	if input.IsDetached {
		flags |= ckalkan.FlagIn2Base64
		flags |= ckalkan.FlagDetachedData

	} else {
		flags |= ckalkan.FlagOutBase64
	}

	if !input.MustCheckCertTime {
		flags |= ckalkan.FlagNoCheckCertTime
	}

	return cli.kc.VerifyXML(signedXML, "", flags)
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
