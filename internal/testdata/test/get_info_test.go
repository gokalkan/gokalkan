package test

import (
	"testing"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func TestX509CertificateGetInfo(t *testing.T) {
	for _, key := range keys {
		t.Run(key.Alias, func(t *testing.T) {
			_, err := cli.X509CertificateGetInfo(key.Cert, ckalkan.CertPropCertCN)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
