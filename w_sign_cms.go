package gokalkan

// SignCMSB64 подписывает данные в base64 и возвращает CMS с подписью и вложенными данными.
func (cli *Client) SignCMSB64(data string, withTSP bool) (signedCMSB64 string, err error) {
	signType := SignBase64
	if withTSP {
		signType |= KCFlagWithTimestamp
	}
	return cli.kc.KCSignData("", data, "", signType)
}

// SignDetachedCMSB64 подписывает данные в base64 и возвращает отделенную подпись.
func (cli *Client) SignDetachedCMSB64(data string, withTSP bool) (signedCMSB64 string, err error) {
	signType := SignBase64 | KCFlagDetachedData
	if withTSP {
		signType |= KCFlagWithTimestamp
	}
	return cli.kc.KCSignData(data, "", "", signType)
}
