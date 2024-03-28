package gokalkan

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// Kalkan - это обертка над методами KalkanCrypt.
type Kalkan interface {
	LoadKeyStore(path, password string) (err error)
	LoadKeyStoreFromBytes(key []byte, password string) (err error)

	Sign(data []byte, isDetached, withTSP bool) (signature []byte, err error)
	SignXML(xml string) (signedXML string, err error)
	SignWSSE(xml, id string) (signedXML string, err error)

	Verify(signature []byte) (string, error)
	VerifyXML(signedXML string) (string, error)
	VerifyDetached(signature, data []byte) (string, error)

	GetCertFromCMS(cms []byte, signID int) (string, error)
	GetCertFromXML(xml string, signID int) (string, error)

	ValidateCert(cert string) (string, error)
	ValidateCertOCSP(cert string, url ...string) (string, error)

	X509CertificateGetInfo(inCert string, prop int) (string, error)

	HashSHA256(data []byte) ([]byte, error)
	HashGOST95(data []byte) ([]byte, error)
	Close() error
}

var _ Kalkan = (*Client)(nil)

var (
	ErrInit    = errors.New("unable to refer to KC_GetFunctionList")
	ErrHTTPCli = errors.New("http cli error")
)

type Client struct {
	log Logger
	kc  *ckalkan.Client
	o   Options
	mu  sync.Mutex
}

// NewClient возвращает клиента для работы с KC.
func NewClient(opts ...Option) (*Client, error) {
	o := Options{log: defaultLogger}
	o.setDefaults()

	for _, op := range opts {
		op(&o)
	}

	kc, err := ckalkan.NewClient()
	if err != nil {
		return nil, err
	}

	cli := &Client{
		log: o.log,
		kc:  kc,
		o:   o,
	}

	err = cli.kc.Init()
	if err != nil {
		cli.log.Error("kc init error: ", err)
		return nil, fmt.Errorf("%w: %s", ErrInit, err)
	}

	cli.kc.TSASetURL(cli.o.TSP)

	if cli.o.LoadCerts {
		if err := cli.LoadCerts(); err != nil {
			cli.log.Error("load remote CA certs error: ", err)
			return nil, err
		}
	}

	return cli, nil
}
