package gokalkan

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Kalkan - это обертка над методами KalkanCrypt (KC)
type Kalkan interface {
	LoadCertsGOST(ctx context.Context) (err error)
	LoadCertsRSA(ctx context.Context) (err error)
	LoadKeyStore(path, password string) (err error)
	LoadKeyStoreFromBytes(key []byte, password string) (err error)

	SignXML(dataXML string) (signedXML string, err error)
	SignCMSB64(signB64, dataB64 string, withTSP bool) (signedCMSB64 string, err error)
	SignDetachedCMSB64(signB64, dataB64 string, withTSP bool) (signedCMSB64 string, err error)
	SignWSSE(dataXML, id string) (signedXML string, err error)

	VerifyXML(signedXML string) (result string, err error)
	VerifyCMSB64(signedCMSB64 string) (result *VerifiedData, err error)
	VerifyDetachedCMSB64(signedCMSB64, dataB64 string) (result *VerifiedData, err error)
	VerifyCert(cert string, t KCValidateType, path ...string) (result string, err error)

	GetCertInfo(certPEM string) (result *X509RawInfo, err error)
	GetCertKeyUsage(certPEM string) (result KeyUsage, err error)
	GetCertFromCMSB64(signedCMSB64 string) (certPEM string, err error)
	GetCertFromKeyStore() (certPEM string, err error)
	Close() error
}

var _ Kalkan = (*Client)(nil)

var (
	ErrLoadKey = errors.New("load key error")
	ErrInit    = errors.New("unable to refer to KC_GetFunctionList")
	ErrHTTPCli = errors.New("http cli error")
)

type Client struct {
	log Logger
	kc  KC
	o   Options
	c   *http.Client
	mu  sync.Mutex
}

// NewClient возвращает клиента для работы с KC.
func NewClient(opts ...Option) (*Client, error) {
	o := Options{log: defaultLogger}

	o.setDefaults()

	for _, op := range opts {
		op(&o)
	}

	o.log.Debug("---------kalkan-config-------------")
	o.log.Debug("Load CA certs on init: ", o.LoadCACertsOnInit)
	o.log.Debug("Load CRL cache on init: ", o.LoadCRLCacheOnInit)
	o.log.Debug("TSP url: ", o.TSP)
	o.log.Debug("OCSP url: ", o.OCSP)
	o.log.Debug("CRL cache duration: ", o.CRLCacheDuration)
	o.log.Debug("CRL GOST url: ", o.CRLGOST)
	o.log.Debug("CRL RSA url: ", o.CRLRSA)
	o.log.Debug("CA GOST url: ", o.CACertGOST)
	o.log.Debug("CA RSA url: ", o.CACertRSA)
	o.log.Debug("NCA RSA url: ", o.NcaCertGOST)
	o.log.Debug("NCA RSA url: ", o.NcaCertRSA)

	kc, err := NewKCClient()
	if err != nil {
		return nil, err
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxIdleConnsPerHost = 100
	t.DisableKeepAlives = true

	if o.Proxy != nil {
		o.log.Debug("Proxy: ", o.Proxy.Hostname())
		t.Proxy = http.ProxyURL(o.Proxy)
	}

	c := &http.Client{
		Transport: t,
		Timeout:   time.Second * 5,
	}

	cli := &Client{
		log: o.log,
		kc:  kc,
		o:   o,
		c:   c,
	}

	cli.log.Debug("kc init...")

	err = cli.kc.KCInit()
	if err != nil {
		cli.log.Error("kc init error: ", err)
		return nil, fmt.Errorf("%w: %s", ErrInit, err)
	}
	cli.log.Debug("kc init OK")
	cli.log.Debug("setting TSP: ", cli.o.TSP)
	cli.kc.KCTSASetURL(cli.o.TSP)

	if cli.o.Proxy != nil {
		cli.log.Debug("setting proxy: ", cli.o.Proxy.Hostname())

		er := cli.kc.KCSetProxy(KCFlagProxyOn, cli.o.Proxy)
		if er != nil {
			cli.log.Error("setting proxy error: ", er)
		}

		cli.log.Debug("setting proxy OK")
	}

	if cli.o.LoadCACertsOnInit {
		var er error

		cli.log.Debug("loading CA certs RSA...")
		er = cli.LoadCertsRSA(context.Background())
		if er != nil {
			cli.log.Error("load CA certs RSA error: ", er)
		}

		cli.log.Debug("load CA certs RSA OK")
		cli.log.Debug("loading CA certs GOST...")

		er = cli.LoadCertsGOST(context.Background())
		if er != nil {
			cli.log.Error("load CA certs GOST error: ", er)
		}

		cli.log.Debug("load CA certs GOST OK")
	}

	if cli.o.LoadCRLCacheOnInit {
		cli.log.Debug("loading CRL cache...")

		er := cli.LoadCRLCache(context.Background())
		if er != nil {
			cli.log.Error("load CRL cache error: ", er)
		}

		cli.log.Debug("load CRL cache OK")
	}

	return cli, nil
}
