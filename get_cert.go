package gokalkan

import (
	"crypto/x509"
	"encoding/base64"
	"github.com/gokalkan/gokalkan/ckalkan"
	"strings"
)

var errUserCertNotFound = ("ERROR 0x8f0001b:")

// GetCertFromCMS обеспечивает получение сертификата из CMS.
func (cli *Client) GetCertFromCMS(cms []byte) ([]*x509.Certificate, error) {
	var (
		certs []*x509.Certificate
		flags = ckalkan.FlagInBase64 | ckalkan.FlagSignCMS | ckalkan.FlagOutBase64
	)

	signID := 1

	for {
		certBase64, err := cli.kc.GetCertFromCMS(base64.StdEncoding.EncodeToString(cms), signID, flags)
		if err != nil {
			if strings.Contains(err.Error(), errUserCertNotFound) && len(certs) != 0 {
				break
			}
			return nil, err
		}

		certBytes, err := base64.StdEncoding.DecodeString(certBase64)
		if err != nil {
			return nil, err
		}

		cert, err := x509.ParseCertificate(certBytes)
		if err != nil {
			return nil, err
		}

		certs = append(certs, cert)
		signID++
	}

	return certs, nil
}

// GetCertFromXML обеспечивает получение сертификата из XML.
func (cli *Client) GetCertFromXML(xml string) ([]*x509.Certificate, error) {
	var signID = 1

	certBase64, err := cli.kc.GetCertFromXML(xml, signID)
	if err != nil {
		return nil, err
	}

	certBytes, err := base64.StdEncoding.DecodeString(string(certBase64))
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	return []*x509.Certificate{cert}, nil
}
