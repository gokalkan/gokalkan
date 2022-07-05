package wtest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gokalkan/gokalkan"
)

func TestGetCertFromStore(t *testing.T) {
	gotCrt, err := cli.GetCertFromKeyStore()
	if err != nil {
		t.Fatal(err)
	}

	info, err := cli.GetCertInfo(gotCrt)
	if err != nil {
		t.Fatal(key.Alias, err)
	}

	json.MarshalIndent(info, "", "\t")

	// t.Log(key.Alias, string(b))

	if gotCrt != key.Cert {
		fmt.Printf("\ngot cert: \n<<<%s>>>\n", gotCrt)
		fmt.Printf("\nexp cert: \n<<<%s>>>\n", key.Cert)
		t.Fatal(key.Alias, " cert mismatch")
	}
}

func TestKCGetCertFromCMS(t *testing.T) {
	gotCert, err := cli.GetCertFromCMSB64(testCMS, 1)
	if err != nil {
		t.Fatal(err)
	}

	if gotCert != testWantCert {
		fmt.Printf("\ngot cert: <<<%s>>>\n", gotCert)
		fmt.Printf("\nexp cert: <<<%s>>>\n", testWantCert)
		t.Fatal("cert mismatch")
	}
}

func TestGetCertInfo(t *testing.T) {
	res, err := cli.GetCertInfo(testWantCert)
	if err != nil {
		t.Fatal(err)
	}

	json.MarshalIndent(res, "", "\t")
	// t.Log(string(b))

	info, err := res.GetX509()
	if err != nil {
		t.Fatal(err)
	}

	json.MarshalIndent(info, "", "\t")
	// t.Log(string(b))

	if info.KeyUsage != gokalkan.KeyUsageSign {
		t.Fatal("mismatch")
	}
}

const (
	testCMS      = `MIIFXAYJKoZIhvcNAQcCoIIFTTCCBUkCAQExDjAMBggqgw4DCgEDAQUAMBMGCSqGSIb3DQEHAaAGBAR0ZXN0oIIDxjCCA8IwggNsoAMCAQICFGzYxuhpvR673kvbfwpus6mTb3e7MA0GCSqDDgMKAQEBAgUAMC4xCzAJBgNVBAYTAktaMR8wHQYDVQQDDBbSsNCa0J4gMy4wIChHT1NUIFRFU1QpMB4XDTIxMDExODEzMTAzM1oXDTIyMDExODEzMTAzM1owga0xHjAcBgNVBAMMFdCi0JXQodCi0J7QkiDQotCV0KHQojEVMBMGA1UEBAwM0KLQldCh0KLQntCSMRgwFgYDVQQFEw9JSU4xMjM0NTY3ODkwMTExCzAJBgNVBAYTAktaMRgwFgYDVQQKDA/QkNCeICLQotCV0KHQoiIxGDAWBgNVBAsMD0JJTjEyMzQ1Njc4OTAyMTEZMBcGA1UEKgwQ0KLQldCh0KLQntCS0JjQpzBsMCUGCSqDDgMKAQEBATAYBgoqgw4DCgEBAQEBBgoqgw4DCgEDAQEAA0MABED/GKU23P7GZC40M7I18/wkaPP5QIXz5l14g9Tqt1xGW4OWUpZ+rkdNeDjycRc8PnpzFB83+mZdfRMr1EA9+l+eo4IB0DCCAcwwDgYDVR0PAQH/BAQDAgbAMCgGA1UdJQQhMB8GCCsGAQUFBwMEBggqgw4DAwQBAgYJKoMOAwMEAQIBMB8GA1UdIwQYMBaAFAe+0hxgCERWsCcV6Fc3Q0e711RQMB0GA1UdDgQWBBRpwj9u4GoXpTfvuQtkfUf/KjlgMDBeBgNVHSAEVzBVMFMGByqDDgMDAgEwSDAhBggrBgEFBQcCARYVaHR0cDovL3BraS5nb3Yua3ovY3BzMCMGCCsGAQUFBwICMBcMFWh0dHA6Ly9wa2kuZ292Lmt6L2NwczA8BgNVHR8ENTAzMDGgL6AthitodHRwOi8vdGVzdC5wa2kuZ292Lmt6L2NybC9uY2FfcnNhX3Rlc3QuY3JsMD4GA1UdLgQ3MDUwM6AxoC+GLWh0dHA6Ly90ZXN0LnBraS5nb3Yua3ovY3JsL25jYV9kX3JzYV90ZXN0LmNybDByBggrBgEFBQcBAQRmMGQwOAYIKwYBBQUHMAKGLGh0dHA6Ly90ZXN0LnBraS5nb3Yua3ovY2VydC9uY2FfcnNhX3Rlc3QuY2VyMCgGCCsGAQUFBzABhhxodHRwOi8vdGVzdC5wa2kuZ292Lmt6L29jc3AvMA0GCSqDDgMKAQEBAgUAA0EAdp3K4AgmME+aEA3Tywcms2sRtyKeQULLaGdYIA+5OhD1Vb9i0R2TuT2R5qi6piRqHv8rb3j5mSGzcqGV95M8pzGCAVMwggFPAgEBMEYwLjELMAkGA1UEBhMCS1oxHzAdBgNVBAMMFtKw0JrQniAzLjAgKEdPU1QgVEVTVCkCFGzYxuhpvR673kvbfwpus6mTb3e7MAwGCCqDDgMKAQMBBQCggaIwGAYJKoZIhvcNAQkDMQsGCSqGSIb3DQEHATAcBgkqhkiG9w0BCQUxDxcNMjIwNTEzMDk1MDMyWjAvBgkqhkiG9w0BCQQxIgQgpuGs3QzH4A0CuQvMsuIYkiidHpP2Irh2DLDgdt7x9CswNwYLKoZIhvcNAQkQAi8xKDAmMCQwIgQgsDrRxPeR/R9RmHlRyf07WB+sRjfKOH+5FXWz2E/pjHMwDQYJKoMOAwoBAQEBBQAEQOCHPqpZDwEajFfJ4Rh8uZwAenavom8J1FJc9RdL2IZ/qMTzn+P14p4w2oSFACQSVhnWYULUz0aDoHyFJtAmeZUA`
	testWantCert = `-----BEGIN CERTIFICATE-----
MIIDwjCCA2ygAwIBAgIUbNjG6Gm9HrveS9t/Cm6zqZNvd7swDQYJKoMOAwoBAQEC
BQAwLjELMAkGA1UEBhMCS1oxHzAdBgNVBAMMFtKw0JrQniAzLjAgKEdPU1QgVEVT
VCkwHhcNMjEwMTE4MTMxMDMzWhcNMjIwMTE4MTMxMDMzWjCBrTEeMBwGA1UEAwwV
0KLQldCh0KLQntCSINCi0JXQodCiMRUwEwYDVQQEDAzQotCV0KHQotCe0JIxGDAW
BgNVBAUTD0lJTjEyMzQ1Njc4OTAxMTELMAkGA1UEBhMCS1oxGDAWBgNVBAoMD9CQ
0J4gItCi0JXQodCiIjEYMBYGA1UECwwPQklOMTIzNDU2Nzg5MDIxMRkwFwYDVQQq
DBDQotCV0KHQotCe0JLQmNCnMGwwJQYJKoMOAwoBAQEBMBgGCiqDDgMKAQEBAQEG
CiqDDgMKAQMBAQADQwAEQP8YpTbc/sZkLjQzsjXz/CRo8/lAhfPmXXiD1Oq3XEZb
g5ZSln6uR014OPJxFzw+enMUHzf6Zl19EyvUQD36X56jggHQMIIBzDAOBgNVHQ8B
Af8EBAMCBsAwKAYDVR0lBCEwHwYIKwYBBQUHAwQGCCqDDgMDBAECBgkqgw4DAwQB
AgEwHwYDVR0jBBgwFoAUB77SHGAIRFawJxXoVzdDR7vXVFAwHQYDVR0OBBYEFGnC
P27gahelN++5C2R9R/8qOWAwMF4GA1UdIARXMFUwUwYHKoMOAwMCATBIMCEGCCsG
AQUFBwIBFhVodHRwOi8vcGtpLmdvdi5rei9jcHMwIwYIKwYBBQUHAgIwFwwVaHR0
cDovL3BraS5nb3Yua3ovY3BzMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly90ZXN0
LnBraS5nb3Yua3ovY3JsL25jYV9yc2FfdGVzdC5jcmwwPgYDVR0uBDcwNTAzoDGg
L4YtaHR0cDovL3Rlc3QucGtpLmdvdi5rei9jcmwvbmNhX2RfcnNhX3Rlc3QuY3Js
MHIGCCsGAQUFBwEBBGYwZDA4BggrBgEFBQcwAoYsaHR0cDovL3Rlc3QucGtpLmdv
di5rei9jZXJ0L25jYV9yc2FfdGVzdC5jZXIwKAYIKwYBBQUHMAGGHGh0dHA6Ly90
ZXN0LnBraS5nb3Yua3ovb2NzcC8wDQYJKoMOAwoBAQECBQADQQB2ncrgCCYwT5oQ
DdPLByazaxG3Ip5BQstoZ1ggD7k6EPVVv2LRHZO5PZHmqLqmJGoe/ytvePmZIbNy
oZX3kzyn
-----END CERTIFICATE-----
`
)
