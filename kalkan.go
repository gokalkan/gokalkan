package kalkan

// Kalkan ...
type Kalkan interface {
	Init() error
	GetLastErrorString() string
	LoadKeyStore(password, containerPath string) error
	X509ExportCertificateFromStore() (string, error)
	VerifyData(data string) (*VerifiedData, error)
	Close() error
	SignXML(data string) (string, error)
	VerifyXML(xml string) (string, error)
}
