package gokalkan

func (cli *Client) GetCertFromKeyStore() (certPEM string, err error) {
	return cli.kc.KCX509ExportCertificateFromStore("")
}
