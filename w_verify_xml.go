package gokalkan

// VerifyXML обеспечивает проверку подписи данных в формате XML.
func (cli *Client) VerifyXML(signedXML string) (result string, err error) {
	return cli.kc.KCVerifyXML(signedXML, "", 0)
}
