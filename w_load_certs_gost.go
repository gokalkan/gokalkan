package gokalkan

import "context"

func (cli *Client) LoadCertsGOST(ctx context.Context) (err error) {
	err = cli.LoadCertFromURL(ctx, cli.o.CACertGOST, KCCertTypeCA)
	if err != nil {
		return err
	}

	err = cli.LoadCertFromURL(ctx, cli.o.NcaCertGOST, KCCertTypeIntermediate)
	if err != nil {
		return err
	}

	return nil
}
