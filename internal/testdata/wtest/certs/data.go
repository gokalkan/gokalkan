package certs

import (
	_ "embed"

	"github.com/gokalkan/gokalkan/internal/testdata"
)

//go:embed gost1.p12
var testKeyGOST1 []byte

//go:embed gost1.cer
var testCertGOST1 string

//go:embed gost2.p12
var testKeyGOST2 []byte

//go:embed gost2.cer
var testCertGOST2 string

var TestKeyGOST1 = testdata.Key{
	Path:     "./internal/testdata/wtest/certs/gost1.p12",
	Password: "Qwerty12",
	Key:      testKeyGOST1,
	Alias:    "wtestgost1",
	Valid:    false,
	Cert:     testCertGOST1,
}

var TestKeyGOST2 = testdata.Key{
	Path:     "./internal/testdata/wtest/certs/gost2.p12",
	Password: "Qwerty12",
	Key:      testKeyGOST2,
	Alias:    "wtestgost2",
	Valid:    false,
	Cert:     testCertGOST2,
}
