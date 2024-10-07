package gokalkan

import (
	"encoding/base64"
	"github.com/gokalkan/gokalkan/ckalkan"
	"time"
)

func (cli *Client) GetTimeFromSig(signature []byte, signID int) (time.Time, error) {
	signatureB64 := base64.StdEncoding.EncodeToString(signature)

	return cli.kc.GetTimeFromSig(signatureB64, signID, ckalkan.FlagInBase64)
}
