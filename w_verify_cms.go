package gokalkan

// VerifyCMSB64 обеспечивает проверку подписи CMS в base64.
func (cli *Client) VerifyCMSB64(signedCMSB64 string) (result *VerifiedData, err error) {
	return cli.kc.KCVerifyData(signedCMSB64, "", SignBase64)
}
