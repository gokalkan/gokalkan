package gokalkan

// GetCertFromCMSB64 обеспечивает получение сертификата из CMS base64.
func (cli *Client) GetCertFromCMSB64(cmsB64 string, signID int) (certPEM string, err error) {
	flags := KCFlagInBase64 | KCFlagOutPEM
	return cli.kc.KCGetCertFromCMS(cmsB64, signID, flags)
}
