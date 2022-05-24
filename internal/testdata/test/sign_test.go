package test

import (
	"testing"

	"github.com/gokalkan/gokalkan"
)

func TestKCSignData(t *testing.T) {
	data := "dGVzdA=="

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(tt *testing.T) {
			gotSignedData, err := cli.KCSignData(data, key.Alias, gokalkan.SignBase64)
			if err != nil {
				tt.Fatalf("%s: %s", key.Alias, err)
			}

			tt.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			gotVerifyResult, err := cli.KCVerifyData(gotSignedData, key.Alias, gokalkan.SignBase64)
			if err != nil {
				tt.Fatal(err)
			}

			t.Log(gotVerifyResult)
		})
	}
}

func TestKCSignXML(t *testing.T) {
	data := "<test>data</test>"

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(tt *testing.T) {
			gotSignedData, err := cli.KCSignXML(data, key.Alias, 0, "", "", "")
			if err != nil {
				tt.Fatalf("%s: %s", key.Alias, err)
			}

			tt.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			gotVerifyResult, err := cli.KCVerifyXML(gotSignedData, key.Alias, 0)
			if err != nil {
				tt.Fatal(err)
			}

			t.Log(gotVerifyResult)
		})
	}
}

func TestKCSignWSEE(t *testing.T) {
	data := gokalkan.WrapWithWSSESoapEnvelope("<test>data</test>", "12345")

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(tt *testing.T) {
			gotSignedData, err := cli.KCSignWSSE(data, key.Alias, 0, "")
			if err != nil {
				tt.Fatalf("%s: %s", key.Alias, err)
			}

			tt.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			//verifyResult, err := cli.KCVerifyXML(gotSignedData, key.Alias, 0)
			//if err != nil {
			//	tt.Fatal(err)
			//}
			//
			//tt.Log(verifyResult)
		})
	}
}
