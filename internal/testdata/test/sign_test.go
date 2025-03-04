package test

import (
	"testing"

	"github.com/gokalkan/gokalkan"
	"github.com/gokalkan/gokalkan/ckalkan"
)

func TestSignData(t *testing.T) {
	data := "dGVzdA=="

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			signBase64 := ckalkan.FlagSignCMS | ckalkan.FlagInBase64 | ckalkan.FlagOutBase64

			gotSignedData, err := cli.SignData("", data, key.Alias, signBase64)
			if err != nil {
				t.Fatalf("%s: %s", key.Alias, err)
			}

			t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			gotVerifyResult, err := cli.VerifyData(gotSignedData, data, key.Alias, signBase64)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(gotVerifyResult)
		})
	}
}

func TestSignXML(t *testing.T) {
	data := "<test>data</test>"

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			gotSignedData, err := cli.SignXML(data, key.Alias, 0, "", "", "")
			if err != nil {
				t.Fatalf("%s: %s", key.Alias, err)
			}

			t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			gotVerifyResult, err := cli.VerifyXML(gotSignedData, key.Alias, 0)
			if err != nil {
				t.Fatal(err)
			}

			t.Log(gotVerifyResult)
		})
	}
}

func TestSignWSEE(t *testing.T) {
	data := gokalkan.WrapWithWSSESoapEnvelope("<test>data</test>", "id-BEFF7CB55C69AB1BB514762482966309")

	for _, key := range keys {
		key := key

		t.Run(key.Alias, func(t *testing.T) {
			gotSignedData, err := cli.SignWSSE(data, key.Alias, 0, "")
			if err != nil {
				t.Fatalf("%s: %s", key.Alias, err)
			}

			t.Logf("\n%s: signed data: '%s'\n", key.Alias, gotSignedData)

			// verifyResult, err := cli.VerifyXML(gotSignedData, key.Alias, 0)
			// if err != nil {
			// 	tt.Fatal(err)
			// }

			// tt.Log(verifyResult)
		})
	}
}
