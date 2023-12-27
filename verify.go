package gokalkan

import (
	"encoding/base64"

	"github.com/gokalkan/gokalkan/types"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// Verify обеспечивает проверку подписи CMS в base64.
func (cli *Client) Verify(input *types.VerifyInput) (string, error) {
	var (
		flags             ckalkan.Flag
		signatureAsBase64 string
		dataAsBase64      string
	)

	signatureAsBase64 = base64.StdEncoding.EncodeToString(input.SignatureBytes)
	flags = ckalkan.FlagSignCMS | ckalkan.FlagInBase64

	if input.IsDetached {
		flags |= ckalkan.FlagIn2Base64
		flags |= ckalkan.FlagDetachedData

		dataAsBase64 = base64.StdEncoding.EncodeToString(input.DataBytes)
	} else {
		flags |= ckalkan.FlagOutBase64
	}

	if !input.MustCheckCertTime {
		flags |= ckalkan.FlagNoCheckCertTime
	}

	resp, err := cli.kc.VerifyData(signatureAsBase64, dataAsBase64, "", flags)
	if err != nil {
		return "", err
	}

	return string(resp.Info), err
}

// VerifyXML обеспечивает проверку подписи данных в формате XML.
func (cli *Client) VerifyXML(signedXML string, mustCheckCertTime bool) (result string, err error) {
	var flags ckalkan.Flag

	if mustCheckCertTime {
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
