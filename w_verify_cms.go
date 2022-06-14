package gokalkan

// VerifyCMSB64 обеспечивает проверку подписи CMS в base64.
func (cli *Client) VerifyCMSB64(signedCMSB64 string) (result *VerifiedData, err error) {
	flags := KCFlagSignCMS | KCFlagInBase64 | KCFlagOutBase64
	return cli.kc.KCVerifyData(signedCMSB64, "", "", flags)
}

// VerifyDetachedCMSB64 обеспечивает проверку отделенной подписи CMS (detached signature) в base64.
func (cli *Client) VerifyDetachedCMSB64(signedCMSB64, dataB64 string) (result *VerifiedData, err error) {
	flags := KCFlagSignCMS | KCFlagInBase64 | KCFlagIn2Base64 | KCFlagDetachedData
	return cli.kc.KCVerifyData(signedCMSB64, dataB64, "", flags)
}
