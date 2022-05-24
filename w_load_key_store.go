package gokalkan

// LoadKeyStore загружает PKCS12.
func (cli *Client) LoadKeyStore(path, password string) error {
	return cli.kc.KCLoadKeyStore(password, path, KCStoreTypePKCS12, "")
}
