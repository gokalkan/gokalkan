package gokalkan

import "context"

func (cli *Client) LoadCertsRSA(ctx context.Context) (err error) {
	err = cli.LoadCertFromURL(ctx, cli.o.CACertRSA, KCCertTypeCA)
	if err != nil {
		return err
	}

	err = cli.LoadCertFromURL(ctx, cli.o.NcaCertRSA, KCCertTypeIntermediate)
	if err != nil {
		return err
	}

	return nil
}
