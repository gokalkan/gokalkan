package gokalkan

import (
	"net/url"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func (cli *Client) SetProxyOn(proxyURL string) error {
	url, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	flag := ckalkan.FlagProxyOn
	return cli.kc.SetProxy(flag, url)
}

func (cli *Client) SetProxyOff(proxyURL string) error {
	url, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	flag := ckalkan.FlagProxyOff
	return cli.kc.SetProxy(flag, url)
}
