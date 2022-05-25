package gokalkan

import (
	"context"
	"os"
	"time"
)

// LoadCRLCache загружает CRL в кэш. Если время жизни кэша истекло, удаляет старые файлы и загружает новые.
// Время жизни можно устанавливать в опциях при создании клиента WithCRLCacheDuration.
func (cli *Client) LoadCRLCache(ctx context.Context) (err error) {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	if cli.o.CRLCacheValidUntil.After(time.Now()) {
		return nil
	}

	cli.log.Debug("CRL cache updating crl files")

	crlGOST, err := cli.downloadCRL(ctx, cli.o.CRLGOST)
	if err != nil {
		cli.log.Error("CRL cache GOST download error: ", err)
		return err
	}

	crlRSA, err := cli.downloadCRL(ctx, cli.o.CRLRSA)
	if err != nil {
		cli.log.Error("CRL cache RSA download error: ", err)
		return err
	}

	crlDeltaGOST, err := cli.downloadCRL(ctx, cli.o.CRLDeltaGOST)
	if err != nil {
		cli.log.Error("CRL cache delta GOST download error: ", err)
		return err
	}

	crlDeltaRSA, err := cli.downloadCRL(ctx, cli.o.CRLDeltaRSA)
	if err != nil {
		cli.log.Error("CRL cache delta RSA download error: ", err)
		return err
	}

	cli.log.Debug("CRL cache updating OK")

	oldCrlGOST := cli.o.crlCache.CRLCachePathGOST
	oldCrlRSA := cli.o.crlCache.CRLCachePathRSA
	oldCrlDeltaGOST := cli.o.crlCache.CRLCachePathDeltaGOST
	oldCrlDeltaRSA := cli.o.crlCache.CRLCachePathDeltaRSA

	defer func() {
		if oldCrlGOST != "" {
			er := os.Remove(oldCrlGOST)
			if er != nil {
				cli.log.Error("CRL cache GOST: remove old cache file error: ", er)
			}
		}

		if oldCrlRSA != "" {
			er := os.Remove(oldCrlRSA)
			if er != nil {
				cli.log.Error("CRL cache RSA: remove old cache file error: ", er)
			}
		}

		if oldCrlDeltaGOST != "" {
			er := os.Remove(oldCrlDeltaGOST)
			if er != nil {
				cli.log.Error("CRL cache delta GOST: remove old cache file error: ", er)
			}
		}

		if oldCrlDeltaRSA != "" {
			er := os.Remove(oldCrlDeltaRSA)
			if er != nil {
				cli.log.Error("CRL cache delta RSA: remove old cache file error: ", er)
			}
		}
	}()

	cli.o.crlCache.CRLCachePathGOST = crlGOST
	cli.o.crlCache.CRLCachePathRSA = crlRSA
	cli.o.crlCache.CRLCachePathDeltaGOST = crlDeltaGOST
	cli.o.crlCache.CRLCachePathDeltaRSA = crlDeltaRSA

	cli.log.Debug("CRL cache GOST: set new cache file: ", crlGOST)
	cli.log.Debug("CRL cache RSA: set new cache file: ", crlRSA)
	cli.log.Debug("CRL cache delta GOST: set new cache file: ", crlDeltaGOST)
	cli.log.Debug("CRL cache delta RSA: set new cache file: ", crlDeltaRSA)

	cli.o.crlCache.CRLCacheValidUntil = time.Now().Add(cli.o.crlCache.CRLCacheDuration)

	cli.log.Debug("cache valid until: ", cli.o.crlCache.CRLCacheValidUntil)

	return nil
}

func (cli *Client) downloadCRL(ctx context.Context, url string) (filename string, err error) {
	tmpFile, err := os.CreateTemp("", "crl.cache.*.crl")
	if err != nil {
		return filename, err
	}

	filename = tmpFile.Name()

	defer func() {
		_ = tmpFile.Close()
	}()

	err = cli.download(ctx, url, tmpFile)
	if err != nil {
		return filename, err
	}

	return filename, nil
}
