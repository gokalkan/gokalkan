package wtest

import (
	"fmt"
	"testing"
)

func TestGetCertFromCMS(t *testing.T) {
	data := []byte("test")

	cms, err := cli.Sign(data, false, false)
	if err != nil {
		t.Fatalf("%s: %s", key.Alias, err)
	}

	gotCrt, err := cli.GetCertFromCMS(cms, 1)
	if err != nil {
		t.Fatal(err)
	}

	if gotCrt != key.Cert {
		fmt.Printf("\ngot cert: \n<<<%s>>>\n", gotCrt)
		fmt.Printf("\nexp cert: \n<<<%s>>>\n", key.Cert)
		t.Fatal(key.Alias, " cert mismatch")
	}
}
