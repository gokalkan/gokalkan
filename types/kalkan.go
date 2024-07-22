package types

import (
	"crypto/x509"
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
)

type VerifyInput struct {
	SignatureBytes    []byte
	DataBytes         []byte
	IsDetached        bool
	MustCheckCertTime bool
}
type VerifyXMLInput struct {
	Signature         string
	MustCheckCertTime bool
}

type SignInput struct {
	DataBytes  []byte
	IsDetached bool
	WithTSP    bool
}

type SignXMLInput struct {
	Data    string
	WithTSP bool
}

type SignHashInput struct {
	Algo       ckalkan.HashAlgo
	InHash     []byte
	IsDetached bool
	WithTSP    bool
}
type CertificateInfo struct {
	Subject           string    // Субъект CertPropSubjectDN
	SerialNumber      string    // Серийный номер  CertPropCertCN
	ValidFrom         time.Time // Начало действия сертификата  CertPropNotBefore
	ValidUntil        time.Time // Конец действия сертификата  CertPropNotAfter
	Issuer            string    // Издатель  CertPropIssuerDN
	Policy            string    // Политика регистрационного свидетельства CertPropPoliciesID
	KeyUsage          string    // Использования ключа CertPropKeyUsage
	ExtKeyUsage       string    // Расширенные использования ключа CertPropExtKeyUsage
	AuthKeyID         string    // Идентификатор ключа центра сертификации
	SubjKeyID         string    // Идентификатор ключа субъекта
	AlgorithmSignCert string    // Алгоритм подписи сертификата CertPropSignatureAlg
	PublicKey         string    // Открытого ключа
	OcspUrl           string    // URL-адрес OCSP
	CrlUrl            string    // URL-адрес CRL
	DeltaCrlUrl       string    // URL-адрес delta CRL
	Policies          []string
	KeyUsages         []string
	ExtKeyUsages      []string
}

type ValidateType string

const (
	ValidateOCSP    ValidateType = "OCSP"
	ValidateCRL     ValidateType = "CRL"
	ValidateNothing ValidateType = "Nothing"
)

type ValidateCertInput struct {
	Certificate   *x509.Certificate
	CheckCertTime bool
	ValidateType  ValidateType
	OCSPUrl       string
	CRLPath       string
}

// Kalkan - это обертка над методами KalkanCrypt.
type Kalkan interface {
	LoadKeyStore(path, password string) (err error)
	LoadKeyStoreFromBytes(key []byte, password string) (err error)
	X509ExportCertificateFromStore() (*x509.Certificate, error)

	Sign(input *SignInput) (signature []byte, err error)
	SignXML(input *SignXMLInput) (string, error)
	SignWSSE(xml, id string) (signedXML string, err error)
	SignHash(input *SignHashInput) (signedHash []byte, err error)

	Verify(input *VerifyInput) (string, error)
	VerifyXML(input *VerifyXMLInput) (result string, err error)
	VerifyDetached(signature, data []byte) (string, error)

	ValidateCert(input *ValidateCertInput) (string, error)

	HashSHA256(data []byte) ([]byte, error)
	HashGOST95(data []byte) ([]byte, error)

	SetProxy(flags ckalkan.Flag, proxyURL string) error

	GetCertFromCMS(cms []byte) ([]*x509.Certificate, error)
	GetCertFromXML(xml string) ([]*x509.Certificate, error)

	X509CertificateGetInfo(*x509.Certificate) (*CertificateInfo, error)
	GetTimeFromSig(cmsDer []byte) (time.Time, error)
	GetSigAlgFromXML(xml string) (string, error)

	Close() error
}
