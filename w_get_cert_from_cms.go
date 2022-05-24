package gokalkan

// GetCertFromCMSB64 обеспечивает получение сертификата из CMS base64.
func (cli *Client) GetCertFromCMSB64(cmsB64 string) (certPEM string, err error) {
	return cli.kc.KCGetCertFromCMS(cmsB64, KCFlagInBase64|KCFlagOutPEM)
}
