package kalkan

// Kalkan - интерфейс с методами KalkanCrypt
type Kalkan interface {
	Init() error
	LoadKeyStore(password, containerPath string) error
	SignXML(data string) (string, error)
	SignWSSE(data string) (string, error)
	SignData(data string) (string, error)
	VerifyXML(xml string) (string, error)
	VerifyData(data string) (*VerifiedData, error)
	X509ExportCertificateFromStore() (string, error)
	GetLastErrorString() string
	Close() error
}
