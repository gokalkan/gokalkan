package gokalkan

import (
	"encoding/pem"
	"github.com/gokalkan/gokalkan/ckalkan"
	kalkanTypes "github.com/gokalkan/gokalkan/types"
)

func (cli *Client) ValidateCert(input *kalkanTypes.ValidateCertInput) (string, error) {
	var flags ckalkan.Flag

	validateType, validatePath := setValidationParams(input)

	certPEM := string(pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: input.Certificate.Raw,
	}))

	if !input.CheckCertTime {
		flags |= ckalkan.FlagNoCheckCertTime
	}

	return cli.kc.X509ValidateCertificate(certPEM, validateType, validatePath, flags)
}

func setValidationParams(input *kalkanTypes.ValidateCertInput) (validateType ckalkan.ValidateType, validatePath string) {
	switch input.ValidateType {
	case kalkanTypes.ValidateOCSP:
		validateType = ckalkan.ValidateTypeOCSP
		validatePath = input.OCSPUrl
	case kalkanTypes.ValidateCRL:
		validateType = ckalkan.ValidateTypeCRL
		validatePath = input.CRLPath
	case kalkanTypes.ValidateNothing:
		validateType = ckalkan.ValidateTypeNothing
		validatePath = ""
	}
	return validateType, validatePath
}
