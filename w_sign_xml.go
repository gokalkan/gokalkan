package gokalkan

// SignXML подписывает данные в формате XML.
func (cli *Client) SignXML(dataXML string) (result string, err error) {
	return cli.kc.KCSignXML(dataXML, "", 0, "", "", "")
}
