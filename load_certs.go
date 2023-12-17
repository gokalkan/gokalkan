package gokalkan

import (
	"bytes"
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

// X509ExportCertificateFromStore экспортирует сертификат из хранилища в формате PEM или в кодировке BASE64
func (cli *Client) X509ExportCertificateFromStore(outputPEM bool) (tokens string, err error) {
	var alias string
	var flags ckalkan.Flag
	if outputPEM {
		flags = ckalkan.FlagOutPEM
	} else {
		flags = ckalkan.FlagOutBase64
	}

	return cli.kc.X509ExportCertificateFromStore(alias, flags)
}
