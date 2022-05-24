package gokalkan

// SignCMSB64 подписывает данные в base64.
func (cli *Client) SignCMSB64(dataB64 string) (signedCMSB64 string, err error) {
	return cli.kc.KCSignData(dataB64, "", SignBase64)
}
