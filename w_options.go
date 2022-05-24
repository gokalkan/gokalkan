package gokalkan

import (
	"net/url"
	"time"
)

func (cli *Client) Options() Options {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	return cli.o
}

type Options struct {
	log          Logger
	TSP          string   `json:"tsp"`            // URL метки времени
	OCSP         string   `json:"ocsp"`           // URL сервиса онлайн проверки статуса сертификата
	Proxy        *url.URL `json:"proxy"`          // URL для прокси
	CACertGOST   string   `json:"ca_cert_gost"`   // URL к сертификату КУЦ
	CACertRSA    string   `json:"ca_cert_rsa"`    // URL к сертификату КУЦ
	NcaCertGOST  string   `json:"nca_cert_gost"`  // URL к сертификату НУЦ
	NcaCertRSA   string   `json:"nca_cert_rsa"`   // URL к сертификату НУЦ
	CRLGOST      string   `json:"crl_gost"`       // URL базового CRL (GOST)
	CRLRSA       string   `json:"crl_rsa"`        // URL базового CRL (RSA)
	CRLDeltaGOST string   `json:"crl_delta_gost"` // URL Дельта CRL (GOST)
	CRLDeltaRSA  string   `json:"crl_delta_rsa"`  // URL Дельта CRL (RSA)
	crlCache
	LoadCRLCacheOnInit bool `json:"load_crl_cache_on_init"`
	LoadCACertsOnInit  bool `json:"load_ca_certs_on_init"`
}

type crlCache struct {
	CRLCachePathGOST      string        `json:"crl_cache_path_gost"`
	CRLCachePathRSA       string        `json:"crl_cache_path_rsa"`
	CRLCachePathDeltaGOST string        `json:"crl_cache_path_delta_gost"`
	CRLCachePathDeltaRSA  string        `json:"crl_cache_path_delta_rsa"`
	CRLCacheValidUntil    time.Time     `json:"crl_cache_valid_until"`
	CRLCacheDuration      time.Duration `json:"crl_cache_duration"`
}

func (o *Options) setDefaults() {
	if o == nil {
		return
	}

	if o.TSP == "" {
		o.TSP = prodTSP
	}

	if o.OCSP == "" {
		o.OCSP = prodOCSP
	}

	if o.CACertGOST == "" {
		o.CACertGOST = prodCACertGOST
	}

	if o.CACertRSA == "" {
		o.CACertRSA = prodCACertRSA
	}

	if o.NcaCertGOST == "" {
		o.NcaCertGOST = prodNcaCertGOST
	}

	if o.NcaCertRSA == "" {
		o.NcaCertRSA = prodNcaCertRSA
	}

	if o.CRLGOST == "" {
		o.CRLGOST = prodCRLGOST
	}

	if o.CRLRSA == "" {
		o.CRLRSA = prodCRLRSA
	}

	if o.CRLDeltaGOST == "" {
		o.CRLDeltaGOST = prodCRLDeltaGOST
	}

	if o.CRLDeltaRSA == "" {
		o.CRLDeltaRSA = prodCRLDeltaRSA
	}

	if o.CRLCacheDuration == 0 {
		o.CRLCacheDuration = prodCRLCacheDuration
	}

	o.LoadCRLCacheOnInit = true
	o.LoadCACertsOnInit = true
}

type Option func(o *Options)

func WithLogger(logger Logger) Option {
	return func(o *Options) {
		o.log = logger
	}
}

func WithTSP(u string) Option {
	return func(o *Options) {
		o.TSP = u
	}
}

func WithProxy(u *url.URL) Option {
	return func(o *Options) {
		o.Proxy = u
	}
}

func WithOCSP(u string) Option {
	return func(o *Options) {
		o.OCSP = u
	}
}

func WithCACertGOST(u string) Option {
	return func(o *Options) {
		o.CACertGOST = u
	}
}

func WithCACertRSA(u string) Option {
	return func(o *Options) {
		o.CACertRSA = u
	}
}

func WithNcaCertGOST(u string) Option {
	return func(o *Options) {
		o.NcaCertGOST = u
	}
}

func WithNcaCertRSA(u string) Option {
	return func(o *Options) {
		o.NcaCertRSA = u
	}
}

func WithCRLGOST(u string) Option {
	return func(o *Options) {
		o.CRLGOST = u
	}
}

func WithCRLRSA(u string) Option {
	return func(o *Options) {
		o.CRLRSA = u
	}
}

func WithCRLDeltaGOST(u string) Option {
	return func(o *Options) {
		o.CRLDeltaGOST = u
	}
}

func WithCRLDeltaRSA(u string) Option {
	return func(o *Options) {
		o.CRLDeltaRSA = u
	}
}

func WithCRLCacheDuration(d time.Duration) Option {
	return func(o *Options) {
		o.CRLCacheDuration = d
	}
}

func WithLoadCRLCacheOnInit(load bool) Option {
	return func(o *Options) {
		o.LoadCRLCacheOnInit = load
	}
}

func WithLoadCACertsOnInit(load bool) Option {
	return func(o *Options) {
		o.LoadCACertsOnInit = load
	}
}

const (
	prodOCSP             = "http://ocsp.pki.gov.kz"
	prodTSP              = "http://tsp.pki.gov.kz:80"
	prodCACertGOST       = "https://pki.gov.kz/cert/root_gost.crt" // КУЦ
	prodCACertRSA        = "https://pki.gov.kz/cert/root_rsa.crt"  // КУЦ
	prodNcaCertGOST      = "https://pki.gov.kz/cert/nca_gost.crt"  // НУЦ
	prodNcaCertRSA       = "https://pki.gov.kz/cert/nca_rsa.crt"   // НУЦ
	prodCRLGOST          = "https://crl.pki.gov.kz/nca_gost.crl"
	prodCRLRSA           = "https://crl.pki.gov.kz/nca_rsa.crl"
	prodCRLDeltaGOST     = "https://crl.pki.gov.kz/nca_d_gost.crl"
	prodCRLDeltaRSA      = "https://crl.pki.gov.kz/nca_d_rsa.crl"
	prodCRLCacheDuration = time.Minute * 60
)

const (
	testOCSP             = "http://test.pki.gov.kz/ocsp/"
	testTSP              = "http://test.pki.gov.kz/tsp/"
	testCACertGOST       = "http://test.pki.gov.kz/cert/root_gost_test.cer" // КУЦ
	testCACertRSA        = "http://test.pki.gov.kz/cert/root_rsa_test.cer"  // КУЦ
	testNcaCertGOST      = "http://test.pki.gov.kz/cert/nca_gost_test.cer"  // НУЦ
	testNcaCertRSA       = "http://test.pki.gov.kz/cert/nca_rsa_test.cer"   // НУЦ
	testCRLGOST          = "http://test.pki.gov.kz/crl/nca_gost_test.crl"
	testCRLRSA           = "http://test.pki.gov.kz/crl/nca_rsa_test.crl"
	testCRLDeltaGOST     = "http://test.pki.gov.kz/crl/nca_d_gost_test.crl"
	testCRLDeltaRSA      = "http://test.pki.gov.kz/crl/nca_d_rsa_test.crl"
	testCRLCacheDuration = time.Second * 5
)

//nolint:gochecknoglobals
var OptsProd = []Option{
	WithLogger(defaultLogger),
	WithTSP(prodTSP),
	WithOCSP(prodOCSP),
	WithCACertGOST(prodCACertGOST),
	WithCACertRSA(prodCACertRSA),
	WithNcaCertGOST(prodNcaCertGOST),
	WithNcaCertRSA(prodNcaCertRSA),
	WithCRLGOST(prodCRLGOST),
	WithCRLRSA(prodCRLRSA),
	WithCRLDeltaGOST(prodCRLDeltaGOST),
	WithCRLDeltaRSA(prodCRLDeltaRSA),
	WithCRLCacheDuration(prodCRLCacheDuration),
	WithLoadCRLCacheOnInit(true),
	WithLoadCACertsOnInit(true),
}

//nolint:gochecknoglobals
var OptsTest = []Option{
	WithLogger(defaultLogger),
	WithTSP(testTSP),
	WithOCSP(testOCSP),
	WithCACertGOST(testCACertGOST),
	WithCACertRSA(testCACertRSA),
	WithNcaCertGOST(testNcaCertGOST),
	WithNcaCertRSA(testNcaCertRSA),
	WithCRLGOST(testCRLGOST),
	WithCRLRSA(testCRLRSA),
	WithCRLDeltaGOST(testCRLDeltaGOST),
	WithCRLDeltaRSA(testCRLDeltaRSA),
	WithCRLCacheDuration(testCRLCacheDuration),
	WithLoadCRLCacheOnInit(true),
	WithLoadCACertsOnInit(true),
}
