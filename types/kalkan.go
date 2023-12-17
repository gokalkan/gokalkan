package types

import "github.com/gokalkan/gokalkan/ckalkan"

type VerifyInput struct {
	SignatureBytes    []byte
	DataBytes         []byte
	IsDetached        bool
	MustCheckCertTime bool
}

// Kalkan - это обертка над методами KalkanCrypt.
type Kalkan interface {
	LoadKeyStore(path, password string) (err error)
	LoadKeyStoreFromBytes(key []byte, password string) (err error)

	Sign(data []byte, isDetached, withTSP bool) (signature []byte, err error)
	SignXML(xml string, withTSP bool) (signedXML string, err error)
	SignWSSE(xml, id string) (signedXML string, err error)
	SignHash(algo ckalkan.HashAlgo, inHash []byte, isDetached, withTSP bool) (signedHash []byte, err error)

	Verify(input *VerifyInput) (string, error)
	VerifyXML(input *VerifyInput) (result string, err error)

	ValidateCert(cert string) (string, error)
	ValidateCertOCSP(cert string, url ...string) (string, error)

	HashSHA256(data []byte) ([]byte, error)
	HashGOST95(data []byte) ([]byte, error)
	Close() error
}
