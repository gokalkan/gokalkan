package gokalkan

// HashData хеширует данные
func (cli *Client) HashData(algo KCHashAlgo, dataB64 string) (hashedB64 string, err error) {
	flags := HashBase64
	return cli.kc.KCHashData(algo, dataB64, flags)
}
