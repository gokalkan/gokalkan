package gokalkan

// SignCMSB64 подписывает данные в base64.
func (cli *Client) SignCMSB64(dataB64 string, withTSP bool) (signedCMSB64 string, err error) {
	signType := SignBase64
	if withTSP {
		signType = SignBase64WithTSP
	}
	return cli.kc.KCSignData(dataB64, "", signType)
}
