package gokalkan

// VerifyCMSB64 обеспечивает проверку подписи CMS в base64.
func (cli *Client) VerifyCMSB64(signedCMSB64, dataB64 string) (result *VerifiedData, err error) {
	flags := KCFlagSignCMS | KCFlagInBase64
	if len(dataB64) != 0 {
		flags |= KCFlagDetachedData
	}
	return cli.kc.KCVerifyData(signedCMSB64, dataB64, "", flags)
}
