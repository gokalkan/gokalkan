package gokalkan

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

// LoadCertFromURL загружает сертификат по url в хранилище.
func (cli *Client) LoadCertFromURL(ctx context.Context, url string, t KCCertType) (err error) {
	tmpCertFile, err := os.CreateTemp("", "tmp.cert.*.crt")
	if err != nil {
		return err
	}

	filename := tmpCertFile.Name()

	defer func() {
		_ = os.Remove(filename)
	}()

	defer func() {
		_ = tmpCertFile.Close()
	}()

	err = cli.download(ctx, url, tmpCertFile)
	if err != nil {
		return err
	}

	err = tmpCertFile.Close()
	if err != nil {
		return err
	}

	err = cli.kc.KCX509LoadCertificateFromFile(filename, t)
	if err != nil {
		return err
	}

	return nil
}

func (cli *Client) download(ctx context.Context, url string, w io.Writer) (err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrHTTPCli, err)
	}

	req.Close = true

	resp, err := cli.c.Do(req)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrHTTPCli, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, resp.Body)

		return fmt.Errorf("%w: bad status: %d", ErrHTTPCli, resp.StatusCode)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrHTTPCli, err)
	}

	return nil
}
