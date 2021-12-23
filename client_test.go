package kalkan

import (
	_ "embed"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	certPath = "test_cert/GOSTKNCA.p12"
	//go:embed test_cert/password.txt
	certPassword string
	cli          Kalkan
)

func TestX509ExportCertificateFromStore(t *testing.T) {
	cert, err := cli.X509ExportCertificateFromStore()
	require.NoError(t, err, "Error on X509ExportCertificateFromStore")
	require.NotEmpty(t, cert, "Cert on X509ExportCertificateFromStore")
}

func TestSignXML(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expError    bool
		expEmptyXML bool
	}{
		{
			name:        "simple",
			input:       "<root>GoKalkan</root>",
			expError:    false,
			expEmptyXML: false,
		},
		{
			name:        "broken xml",
			input:       "<root>GoKalkan</invalid>",
			expError:    true,
			expEmptyXML: true,
		},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			xml, err := cli.SignXML(ts.input)
			expectError(assert.New(t), "SignXML", ts.expError, err)
			expectEmpty(assert.New(t), "SignXML", ts.expEmptyXML, xml)
		})
	}
}

func TestVerifyXML(t *testing.T) {
	tests := []struct {
		name        string
		rawXML      string
		editSigned  func(string) string
		expError    bool
		expEmptyXML bool
	}{
		{
			name:        "simple",
			rawXML:      "<root>GoKalkan</root>",
			editSigned:  func(s string) string { return s },
			expError:    true,
			expEmptyXML: true,
		},
	}
	for _, ts := range tests {
		t.Run(ts.name, func(t *testing.T) {
			xml, err := cli.SignXML(ts.rawXML)
			expErrOnSign := false
			expEmptyXMLOnSign := false
			expectError(assert.New(t), "SignXML", expErrOnSign, err)
			expectEmpty(assert.New(t), "SignXML", expEmptyXMLOnSign, xml)

			data := ts.editSigned(xml)
			serialNum, err := cli.VerifyXML(data)
			expectError(assert.New(t), "VerifyXML", ts.expError, err)
			expectEmpty(assert.New(t), "VerifyXML", ts.expEmptyXML, serialNum)
		})
	}
}

func TestMain(m *testing.M) {
	setupTest()
	code := m.Run()
	tearDownTest()
	os.Exit(code)
}

func setupTest() {
	tmpCli, err := NewClient()
	if err != nil {
		log.Fatal("NewClient", err)
	}
	if err := tmpCli.LoadKeyStore(certPassword, certPath); err != nil {
		log.Fatal("cli.LoadKeyStore", err)
	}
	cli = tmpCli
}

func tearDownTest() {
	if err := cli.Close(); err != nil {
		log.Fatal("cli.Close", err)
	}
}

func expectError(as *assert.Assertions, name string, expErr bool, err error) bool {
	if expErr {
		return as.Error(err, "Expect error on", name)
	}
	return as.NoError(err, "Expect no error on", name)
}

func expectEmpty(as *assert.Assertions, name string, expEmpty bool, object interface{}) bool {
	if expEmpty {
		return as.Empty(object, "Expect empty object on", name)
	}
	return as.NotEmpty(object, "Expect non-empty object on", name)
}
