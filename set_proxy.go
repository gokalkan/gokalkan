package gokalkan

import (
	"net/url"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// Использование прокси сервера.
func (cli *Client) SetProxy(flags ckalkan.Flag, proxyURL string) error {
	url, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	return cli.kc.SetProxy(flags, url)
}
