package gokalkan

import "github.com/gokalkan/gokalkan/ckalkan"

func (cli *Client) X509CertificateGetInfo(inCert string, prop int) (string, error) {
	return cli.kc.X509CertificateGetInfo(inCert, ckalkan.CertProp(prop))
}
