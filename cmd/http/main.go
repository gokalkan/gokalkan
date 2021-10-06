package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/Zulbukharov/kalkan-bind/pkg/bridge"
	"github.com/Zulbukharov/kalkan-bind/pkg/challenge"
	"github.com/Zulbukharov/kalkan-bind/pkg/httpd"
	"github.com/Zulbukharov/kalkan-bind/pkg/settings"
	"github.com/Zulbukharov/kalkan-bind/pkg/storage/memory"
)

func main() {
	conf, err := settings.ParseYAML("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	b, e := bridge.NewKalkanBridge()
	if e != nil {
		fmt.Println("here?", e)
		return
	}
	b.Init()
	defer b.Close()
	b.KCLoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)

	m := memory.NewStorage()
	challengeS := challenge.NewService(m, b)

	// b.X509ExportCertificateFromStore()
	s, rv := b.SignXML(`<company-id>770704034</company-id>`)
	fmt.Println("SignXML", rv, s)

	challengeHandler := httpd.NewChallengeHandler(challengeS)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", conf.Port),
		Handler: httpd.Route(
			challengeHandler,
		),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}

	// challengeS.HandleChallenge(s)
}

// "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<company-id>770704034<ds:Signature xmlns:ds=\"http://www.w3.org/2000/09/xmldsig#\" Id=\"1\">\n<ds:SignedInfo>\n<ds:CanonicalizationMethod Algorithm=\"http://www.w3.org/TR/2001/REC-xml-c14n-20010315\"/>\n<ds:SignatureMethod Algorithm=\"http://www.w3.org/2001/04/xmldsig-more#rsa-sha256\"/>\n<ds:Reference URI=\"\">\n<ds:Transforms>\n<ds:Transform Algorithm=\"http://www.w3.org/2000/09/xmldsig#enveloped-signature\"/>\n<ds:Transform Algorithm=\"http://www.w3.org/TR/2001/REC-xml-c14n-20010315#WithComments\"/>\n</ds:Transforms>\n<ds:DigestMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#sha256\"/>\n<ds:DigestValue>DNQp4hl/PqszLNFCuJdacBzNazvtHP2N8DFEh0+VhUg=</ds:DigestValue>\n</ds:Reference>\n</ds:SignedInfo>\n<ds:SignatureValue>AFB0ADcRc2JQIoRUDjqyKEr2nklJWl/54WcbmzL9Xev1iBFRWP+572biDmWCJcjj\ntJROmkbIOB4LZUEj+9L23VYcV1X9YdljAKrwt+D5k7PoJEUSXenvxU2H4ngNnYAf\n++h4Lfp/EjxcQS4AhHuq5nKUo+YAFvZZHQupYyG+ogrt7gGr8FDp6CKWV3GLzMai\nwphK6+eAg89WOoz2iJzC19xbfWiyZTcYDvu/+mW2TtBh9AtbjBqu/8e9FUYNT50k\nJU4IOa9XUbMQftf6c/U3Qn2Os0bQOA7mlvz0KTgrSfpD3iCmW0qHrkwWZkXYJNKF\nuxjbzisBQXARulCQsFg8xQ==</ds:SignatureValue>\n<ds:KeyInfo>\n<ds:X509Data>\n<ds:X509Certificate>MIIGODCCBCCgAwIBAgIUORdZ5WWNykkBF0F4OObHLNioU+cwDQYJKoZIhvcNAQEL\nBQAwLTELMAkGA1UEBhMCS1oxHjAcBgNVBAMMFdKw0JrQniAzLjAgKFJTQSBURVNU\nKTAeFw0yMDAxMjgwNjIzMDNaFw0yMTAxMjcwNjIzMDNaMIG1MR4wHAYDVQQDDBXQ\notCV0KHQotCe0JIg0KLQldCh0KIxFTATBgNVBAQMDNCi0JXQodCi0J7QkjEYMBYG\nA1UEBRMPSUlOMTIzNDU2Nzg5MDExMQswCQYDVQQGEwJLWjEcMBoGA1UEBwwT0J3Q\no9CgLdCh0KPQm9Ci0JDQnTEcMBoGA1UECAwT0J3Qo9CgLdCh0KPQm9Ci0JDQnTEZ\nMBcGA1UEKgwQ0KLQldCh0KLQntCS0JjQpzCCASIwDQYJKoZIhvcNAQEBBQADggEP\nADCCAQoCggEBAJH9oFXBXM/LwgbFQCPyChq+1WITJvZr6uuPp3Hmj8H5WwKx21Uw\nCNYDfGyodXgtji+VPOWzv0WHKbNgNyFhytMtPpLCpPUR45mElgOsKfLIhCKnfRhJ\nbFVMXXRg5rtLa7yD97u5r5ykKHQ8GOsel7oIWpqETbfj7uyGQXymZxQG7zR5apVS\nEimbmbclxtLd7CtGrS5CuvaQvxvz9PjvJsZgZ5J4gXhgFnAxMbJsq7tJGNCM1uo3\nbOdeYrhEhWdWzm72tGu8m2jpmGO0BQo0otR1h0p9ALlHNlwECXSBxhU0M5bkLLRD\nFNEboul89qifXVDZcmISznjPOuwnsomLVw8CAwEAAaOCAcUwggHBMA4GA1UdDwEB\n/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAgYIKoMOAwMEAQEwHwYDVR0jBBgw\nFoAUpowWM3y46DVnBj5eQVdVoq80UGgwHQYDVR0OBBYEFLLHCp/3pdxZ3gzMPA3d\n40N8wfEtMF4GA1UdIARXMFUwUwYHKoMOAwMCBDBIMCEGCCsGAQUFBwIBFhVodHRw\nOi8vcGtpLmdvdi5rei9jcHMwIwYIKwYBBQUHAgIwFwwVaHR0cDovL3BraS5nb3Yu\na3ovY3BzMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly90ZXN0LnBraS5nb3Yua3ov\nY3JsL25jYV9yc2FfdGVzdC5jcmwwPgYDVR0uBDcwNTAzoDGgL4YtaHR0cDovL3Rl\nc3QucGtpLmdvdi5rei9jcmwvbmNhX2RfcnNhX3Rlc3QuY3JsMHIGCCsGAQUFBwEB\nBGYwZDA4BggrBgEFBQcwAoYsaHR0cDovL3Rlc3QucGtpLmdvdi5rei9jZXJ0L25j\nYV9yc2FfdGVzdC5jZXIwKAYIKwYBBQUHMAGGHGh0dHA6Ly90ZXN0LnBraS5nb3Yu\na3ovb2NzcC8wDQYJKoZIhvcNAQELBQADggIBAAMoo4+pbazfRVXlA8/E8FMKqu7m\nEH4cGu1cTvSUrb8BBkcf8dfJYy0bRdCJrCQMyG0NEN29c8DV2rMuIMu+5BiJgTfp\n7m3YlAc504nMTT7JkNyF+ewk8hfMuvvlVZSOWsXLtvnUTAzc+3wVUquKRkKgx2kf\n6c0nBMZvbBgbx6ejuhDjIxyTNGNlc9vpEn9vNo1u06MtjCW2qEEkCuGYoVhdfCS+\ng2NFXYFQ2YYoUIFpiPzVKL7nqnShNujCgkudVAJwmxWpABPTOEmk3prl2Qnb2oUn\nUYfnKfkeGwMMeNDE8KwCXwUpyWRAZdGgfa711ZC0ppPLv0WzANKuraVd2T41XrRO\nTV3qzdT2vmNvL+Orig5lUoUMZ2dIuK/bEwLiMsh2Ip/6pofmVfAzWCDSh/5gCy06\nX+dRXWmKK89a8I+lyxbmwypWA474Hqt2SOfbna6mfkPJRrvNq+jCVcgSK1rUh+u+\nd5nnfLV4AVgH74Z203XZODIp/NQCOQnn+ggE5Em8qE6Ylq4jL1TzidnE9EISEeTP\nFdDEHjdAsVvuTxjMOZb98pCr2wSyzJk3WzTjz25zdYoN7CkS/gd+yJlmKxszVUzX\nepZwPrcIGWjIQjG0XCz5NagOqoj+0Lj/px5k8cEAND0NPA5T1lHYyMCMoU1NmGFf\nOecuSbFrJVVAPc3O</ds:X509Certificate>\n</ds:X509Data>\n</ds:KeyInfo>\n</ds:Signature></company-id>\n"
