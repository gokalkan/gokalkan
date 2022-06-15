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

	_, err = cli.VerifyCMSB64(gotSignedData)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignXML(t *testing.T) {
	data := "<test>data</test>"

	gotSignedData, err := cli.SignXML(data)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	_, err = cli.VerifyXML(gotSignedData)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSignWSEE(t *testing.T) {
	gotSignedData, err := cli.SignWSSE("<test>data</test>", "id-12345")
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}
	_, err = cli.VerifyXML(gotSignedData)
	if err != nil {
		t.Fatal(err)
	}
}

// func TestSignCMSB64WithTSP(t *testing.T) {
// 	data := "dGVzdA=="

// 	gotSignedData, err := cli.SignCMSB64(data, true)
// 	if err != nil {
// 		t.Fatalf("%s: %s", key.Alias, err)
// 	}

// 	_, err := cli.VerifyCMSB64(gotSignedData)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

func TestSignDetachedCMSB64(t *testing.T) {
	data := "dGVzdA=="

	gotSignedData, err := cli.SignDetachedCMSB64(data, false)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	_, err = cli.VerifyDetachedCMSB64(gotSignedData, data)
	if err != nil {
		t.Fatal(err)
	}
}
