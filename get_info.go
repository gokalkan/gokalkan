package gokalkan

import (
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// GetTimeFromSig получает время подписания из CMS
// Если вы хотите пeредать данные подписи в base64 формате, то установите флаг base64 = true
func (cli *Client) GetTimeFromSig(data string, base64 bool) (time.Time, error) {
	var flag ckalkan.Flag

	flag = ckalkan.FlagInDER

	if base64 {
		flag = ckalkan.FlagInBase64
	}

	return cli.kc.GetTimeFromSig(data, flag, 0)
}
