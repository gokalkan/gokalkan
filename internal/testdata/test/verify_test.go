package test

import (
	gokalkan "github.com/Zulbukharov/GoKalkan"
	"testing"
)

func TestKCX509ValidateCertificateNothing(t *testing.T) {
	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(tt *testing.T) {
			gotResult, err := cli.KCX509ValidateCertificate(key.Cert, gokalkan.KCValidateTypeNothing, "")
			if err != nil {
				tt.Fatal(err)
			}

			tt.Log(key.Alias, gotResult)
		})
	}
}

func TestKCX509ValidateCertificateOCSP(t *testing.T) {
	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(tt *testing.T) {
			gotResult, err := cli.KCX509ValidateCertificate(key.Cert, gokalkan.KCValidateTypeOCSP, "http://test.pki.gov.kz/ocsp/")
			if err != nil {
				tt.Fatal(err)
			}

			tt.Log(key.Alias, gotResult)
		})
	}
}
