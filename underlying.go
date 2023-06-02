package gokalkan

import "github.com/gokalkan/gokalkan/ckalkan"

// Underlying возвращает текущий ckalkan.Client, используемый клиентом gokalkan.
func (cli *Client) Underlying() *ckalkan.Client {
	return cli.kc
}
