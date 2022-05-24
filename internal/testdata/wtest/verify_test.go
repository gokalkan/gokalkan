package wtest

import (
	gokalkan "github.com/Zulbukharov/GoKalkan"
	"testing"
)

func TestVerifyCertNothing(t *testing.T) {
	gotResult, err := cli.VerifyCert(key.Cert, gokalkan.KCValidateTypeNothing)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(key.Alias, gotResult)
}

func TestVerifyCertOCSP(t *testing.T) {
	gotResult, err := cli.VerifyCert(key.Cert, gokalkan.KCValidateTypeOCSP)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(key.Alias, gotResult)
}

func TestVerifyCertCRL(t *testing.T) {
	gotResult, err := cli.VerifyCert(key.Cert, gokalkan.KCValidateTypeCRL)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(key.Alias, gotResult)
}
