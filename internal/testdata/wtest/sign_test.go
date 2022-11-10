package wtest

import (
	"testing"
)

func TestSign(t *testing.T) {
	data := []byte("test")

	gotSignedData, err := cli.Sign(data, false, false)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	_, err = cli.Verify(gotSignedData)
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

func TestSignDetached(t *testing.T) {
	data := []byte("test")

	gotSignedData, err := cli.Sign(data, true, false)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	_, err = cli.VerifyDetached(gotSignedData, data)
	if err != nil {
		t.Fatal(err)
	}
}
