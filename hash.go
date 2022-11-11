package gokalkan

import (
	"encoding/base64"

	"github.com/gokalkan/gokalkan/ckalkan"
)

func (cli *Client) HashSHA256(data []byte) (hashed []byte, err error) {
	flags := ckalkan.FlagInBase64 | ckalkan.FlagOutBase64
	algo := ckalkan.HashAlgoSHA256
	dataB64 := base64.StdEncoding.EncodeToString(data)

	hashedB64, err := cli.kc.HashData(algo, dataB64, flags)
	if err != nil {
		return nil, err
	}

	raw, err := base64.StdEncoding.DecodeString(hashedB64)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (cli *Client) HashGOST95(data []byte) (hashed []byte, err error) {
	flags := ckalkan.FlagInBase64 | ckalkan.FlagOutBase64
	algo := ckalkan.HashAlgoGOST95
	dataB64 := base64.StdEncoding.EncodeToString(data)

	hashedB64, err := cli.kc.HashData(algo, dataB64, flags)
	if err != nil {
		return nil, err
	}

	raw, err := base64.StdEncoding.DecodeString(hashedB64)
	if err != nil {
		return nil, err
	}
	return raw, nil
}
