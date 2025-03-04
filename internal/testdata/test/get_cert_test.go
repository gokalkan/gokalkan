package test

import (
	"fmt"
	"testing"
)

func TestX509ExportCertificateFromStore(t *testing.T) {
	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			gotCrt, err := cli.X509ExportCertificateFromStore(key.Alias)
			if err != nil {
				t.Fatal(err)
			}

			if gotCrt != key.Cert {
				fmt.Printf("\ngot cert: \n<<<%s>>>\n", gotCrt)
				fmt.Printf("\nexp cert: \n<<<%s>>>\n", key.Cert)
				t.Fatal(key.Alias, " cert mismatch")
			}
		})
	}
}
