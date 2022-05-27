package wtest

import (
	"testing"
)

func TestSignCMSB64(t *testing.T) {
	data := "dGVzdA=="

	gotSignedData, err := cli.SignCMSB64(data, false)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

	gotVerifyResult, err := cli.VerifyCMSB64(gotSignedData)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(gotVerifyResult)

}

func TestSignXML(t *testing.T) {
	data := "<test>data</test>"

	gotSignedData, err := cli.SignXML(data)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

	gotVerifyResult, err := cli.VerifyXML(gotSignedData)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(gotVerifyResult)

}

func TestSignWSEE(t *testing.T) {
	gotSignedData, err := cli.SignWSSE("<test>data</test>", "12345")
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

	//verifyResult, err := cli.KCVerifyXML(gotSignedData, key.Alias, 0)
	//if err != nil {
	//	tt.Fatal(err)
	//}
	//
	//tt.Log(verifyResult)
}
