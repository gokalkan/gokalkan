package gokalkan

import "context"

func (cli *Client) LoadCerts(ctx context.Context) (err error) {
	for _, v := range cli.o.Certs {
		err = cli.LoadCertFromURL(ctx, v.URL, v.Type)
		if err != nil {
			return err
		}
	}
	return nil
}
