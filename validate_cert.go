package gokalkan

import (
	"github.com/gokalkan/gokalkan/ckalkan"
)

func (cli *Client) ValidateCert(cert string) (string, error) {
	validateType := ckalkan.ValidateTypeNothing
	validatePath := ""

	var flags ckalkan.Flag

	flags |= ckalkan.FlagNoCheckCertTime

	return cli.kc.X509ValidateCertificate(cert, validateType, validatePath, flags)
}

func (cli *Client) ValidateCertOCSP(cert string, url ...string) (string, error) {
	validateType := ckalkan.ValidateTypeOCSP
	validatePath := cli.o.OCSP

	var flags ckalkan.Flag

	flags |= ckalkan.FlagNoCheckCertTime

	if len(url) > 0 {
		validatePath = url[0]
	}

	return cli.kc.X509ValidateCertificate(cert, validateType, validatePath, flags)

}
