package kalkan

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	certPath = "test_cert/GOSTKNCA.p12"
	//go:embed test_cert/password.txt
	certPassword string
)

type KalkanSuite struct {
	suite.Suite
	cli Kalkan
}

func (ks *KalkanSuite) SetupTest() {
	cli, err := NewClient()
	if err != nil {
		ks.T().Error("NewClient", err)
	}
	if err := cli.LoadKeyStore(certPassword, certPath); err != nil {
		ks.T().Error("cli.LoadKeyStore", err)
	}
	ks.cli = cli
}

func (ks *KalkanSuite) TearDownTest() {
	if err := ks.cli.Close(); err != nil {
		ks.T().Error("cli.Close", err)
	}
}

func (ks *KalkanSuite) TestX509ExportCertificateFromStore() {
	cert, err := ks.cli.X509ExportCertificateFromStore()
	ks.Require().NoError(err, "Error on X509ExportCertificateFromStore")
	ks.Require().NotEmpty(cert, "Cert on X509ExportCertificateFromStore")
}

func (ks *KalkanSuite) TestSignXML() {
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
		ks.T().Run(ts.name, func(t *testing.T) {
			xml, err := ks.cli.SignXML(ts.input)
			expectError(ks.Assert(), "SignXML", ts.expError, err)
			expectEmpty(ks.Assertions, "SignXML", ts.expEmptyXML, xml)
		})
	}
}

func (ks *KalkanSuite) TestVerifyXML() {
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
			expEmptyXML: false,
		},
	}
	for _, ts := range tests {
		ks.T().Run(ts.name, func(t *testing.T) {
			xml, err := ks.cli.SignXML(ts.rawXML)
			expErrOnSign := false
			expEmptyXMLOnSign := false
			expectError(ks.Assert(), "SignXML", expErrOnSign, err)
			expectEmpty(ks.Assertions, "SignXML", expEmptyXMLOnSign, xml)

			data := ts.editSigned(xml)
			serialNum, err := ks.cli.VerifyXML(data)
			expectError(ks.Assert(), "VerifyXML", ts.expError, err)
			expectEmpty(ks.Assert(), "VerifyXML", ts.expEmptyXML, serialNum)
		})
	}
}

func TestKalkan(t *testing.T) {
	suite.Run(t, new(KalkanSuite))
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
