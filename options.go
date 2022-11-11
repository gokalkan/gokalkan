package gokalkan

import (
	"crypto/x509"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func (cli *Client) Options() Options {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	return cli.o
}

type OptionsCert struct {
	Cert *x509.Certificate
	Type ckalkan.CertType
}

type Options struct {
	log       Logger
	TSP       string        `json:"tsp"`   // URL метки времени
	OCSP      string        `json:"ocsp"`  // URL сервиса онлайн проверки статуса сертификата
	Certs     []OptionsCert `json:"certs"` // Корневые сертификатам
	LoadCerts bool          `json:"load_certs"`
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

	o.LoadCerts = true
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

func WithOCSP(u string) Option {
	return func(o *Options) {
		o.OCSP = u
	}
}

//nolint:gochecknoglobals
const (
	prodOCSP = "http://ocsp.pki.gov.kz"
	prodTSP  = "http://tsp.pki.gov.kz:80"

	testOCSP = "http://test.pki.gov.kz/ocsp/"
	testTSP  = "http://test.pki.gov.kz/tsp/"
)

//nolint:gochecknoglobals
var (
	OptsProd = []Option{
		WithLogger(defaultLogger),
		WithTSP(prodTSP),
		WithOCSP(prodOCSP),
	}
	OptsTest = []Option{
		WithLogger(defaultLogger),
		WithTSP(testTSP),
		WithOCSP(testOCSP),
	}
)
