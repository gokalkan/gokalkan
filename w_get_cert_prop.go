package gokalkan

import "strings"

// GetCertKeyUsage возвращает информацию KeyUsage о сертификате.
func (cli *Client) GetCertKeyUsage(certPEM string) (result KeyUsage, err error) {
	res, err := cli.kc.KCX509CertificateGetInfo(certPEM, KCCertPropKeyUsage)
	if err != nil {
		return result, err
	}

	result = parseKeyUsage(strings.TrimPrefix(res, "keyUsage="))

	return result, nil
}

// GetCertProp возвращает свойства о сертификате
func (cli *Client) GetCertProp(certPEM string, prop KCCertProp) (string, error) {
	res, err := cli.getCertProp(certPEM, prop)
	if err != nil {
		return "", err
	}
	res = strings.TrimSpace(res)
	return res, nil
}
