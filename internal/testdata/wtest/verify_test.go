package wtest

import (
	"testing"
)

func TestVerifyCertNothing(t *testing.T) {
	gotResult, err := cli.ValidateCert(key.Cert)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(key.Alias, gotResult)
}

func TestVerifyCertOCSP(t *testing.T) {
	gotResult, err := cli.ValidateCertOCSP(key.Cert)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(key.Alias, gotResult)
}
