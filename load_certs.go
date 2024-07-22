package gokalkan

import (
	"bytes"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/gokalkan/gokalkan/ckalkan"
)

var ErrLoadKey = errors.New("load key error")

func (cli *Client) LoadCerts() error {
	for _, v := range cli.o.Certs {
		if err := cli.LoadCertFromBytes(v.Cert.Raw, v.Type); err != nil {
			return err
		}
	}
	return nil
}

func (cli *Client) LoadCertFromBytes(cert []byte, certType ckalkan.CertType) (err error) {
	tmpCert, err := os.CreateTemp("", "tmp.cert.*.crt")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	filename := tmpCert.Name()

	defer os.Remove(filename)
	defer tmpCert.Close()

	written, err := io.Copy(tmpCert, bytes.NewReader(cert))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	if exp := int64(len(cert)); exp != written {
		return fmt.Errorf("%w: expected %d bytes, but written %d bytes", ErrLoadKey, exp, written)
	}

	return cli.kc.X509LoadCertificateFromFile(filename, certType)
}

// X509ExportCertificateFromStore экспортирует сертификат из хранилища.
func (cli *Client) X509ExportCertificateFromStore() (cert *x509.Certificate, err error) {
	var (
		alias string
	)

	certBase64, err := cli.kc.X509ExportCertificateFromStore(alias, ckalkan.FlagOutBase64)
	if err != nil {
		return nil, err
	}

	certBytes, err := base64.StdEncoding.DecodeString(certBase64)
	if err != nil {
		return nil, err
	}

	cert, err = x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	return cert, nil
}
