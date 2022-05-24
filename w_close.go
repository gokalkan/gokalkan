package gokalkan

import "os"

func (cli *Client) Close() (err error) {
	err = cli.kc.KCClose()
	if err != nil {
		cli.log.Error("kc close error: ", err)
		return err
	}

	cli.log.Debug("CRL cache clearing...")

	if cli.o.crlCache.CRLCachePathGOST != "" {
		er := os.Remove(cli.o.crlCache.CRLCachePathGOST)
		if er != nil {
			cli.log.Error("CRL cache GOST: remove old cache file error: ", er)
		}
	}

	if cli.o.crlCache.CRLCachePathRSA != "" {
		er := os.Remove(cli.o.crlCache.CRLCachePathRSA)
		if er != nil {
			cli.log.Error("CRL cache RSA: remove old cache file error: ", er)
		}
	}

	if cli.o.crlCache.CRLCachePathDeltaGOST != "" {
		er := os.Remove(cli.o.crlCache.CRLCachePathDeltaGOST)
		if er != nil {
			cli.log.Error("CRL cache delta GOST: remove old cache file error: ", er)
		}
	}

	if cli.o.crlCache.CRLCachePathDeltaRSA != "" {
		er := os.Remove(cli.o.crlCache.CRLCachePathDeltaRSA)
		if er != nil {
			cli.log.Error("CRL cache delta RSA: remove old cache file error: ", er)
		}
	}

	return nil
}
