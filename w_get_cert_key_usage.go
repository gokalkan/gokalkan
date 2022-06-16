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

// GetCertKeyUsages возвращает свойства о сертификате
func (cli *Client) GetCertKeyUsages(certPEM string) ([]string, error) {
	res, err := cli.kc.KCX509CertificateGetInfo(certPEM, KCCertPropKeyUsage)
	if err != nil {
		return nil, err
	}

	ku := strings.Split(strings.TrimSpace(res), " ")

	return ku, nil
}
