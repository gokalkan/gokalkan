package test

import (
	"testing"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func TestX509ValidateCertificateNothing(t *testing.T) {
	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			_, err := cli.X509ValidateCertificate(key.Cert, ckalkan.ValidateTypeNothing, "")
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestX509ValidateCertificateOCSP(t *testing.T) {
	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			_, err := cli.X509ValidateCertificate(key.Cert, ckalkan.ValidateTypeOCSP, "http://test.pki.gov.kz/ocsp/")
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
