package gokalkan

import (
	"context"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func WithCert(cert *x509.Certificate, typ ckalkan.CertType) Option {
	return func(o *Options) {
		o.Certs = append(o.Certs, OptionsCert{cert, typ})
	}
}

func WithCerts(c []OptionsCert) Option {
	return func(o *Options) {
		o.Certs = c
	}
}

func WithRemoteProdCerts(ctx context.Context) Option {
	type pair struct {
		url      string
		certType ckalkan.CertType
	}
	pairs := []pair{
		{url: "https://pki.gov.kz/cert/root_gost.crt", certType: ckalkan.CertTypeCA},
		{url: "https://pki.gov.kz/cert/root_rsa.crt", certType: ckalkan.CertTypeCA},
		{url: "https://pki.gov.kz/cert/root_gost2015_2022.cer", certType: ckalkan.CertTypeCA},
		{url: "https://pki.gov.kz/cert/nca_gost.crt", certType: ckalkan.CertTypeIntermediate},
		{url: "https://pki.gov.kz/cert/nca_rsa.crt", certType: ckalkan.CertTypeIntermediate},
		{url: "https://pki.gov.kz/cert/nca_gost2015.cer", certType: ckalkan.CertTypeIntermediate},
	}

	return func(o *Options) {
		for _, p := range pairs {
			bytes, err := download(ctx, p.url)
			if err != nil {
				panic(err)
			}
			cert, err := x509.ParseCertificate(bytes)
			if err != nil {
				panic(err)
			}
			o.Certs = append(o.Certs, OptionsCert{Cert: cert, Type: p.certType})
		}
	}
}

func WithRemoteTestCerts(ctx context.Context) Option {
	type pair struct {
		url      string
		certType ckalkan.CertType
	}
	pairs := []pair{
		{url: "http://test.pki.gov.kz/cert/root_gost_test.cer", certType: ckalkan.CertTypeCA},
		{url: "http://test.pki.gov.kz/cert/root_rsa_test.cer", certType: ckalkan.CertTypeCA},
		{url: "http://test.pki.gov.kz/cert/nca_gost_test.cer", certType: ckalkan.CertTypeIntermediate},
		{url: "http://test.pki.gov.kz/cert/nca_rsa_test.cer", certType: ckalkan.CertTypeIntermediate},
	}

	return func(o *Options) {
		for _, p := range pairs {
			bytes, err := download(ctx, p.url)
			if err != nil {
				panic(err)
			}
			cert, err := x509.ParseCertificate(bytes)
			if err != nil {
				panic(err)
			}
			o.Certs = append(o.Certs, OptionsCert{Cert: cert, Type: p.certType})
		}
	}
}

// LoadCertFromURL загружает сертификат по url в хранилище.
func download(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(io.Discard, resp.Body) //nolint:errcheck
		return nil, fmt.Errorf("%w: bad status: %d", ErrHTTPCli, resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
