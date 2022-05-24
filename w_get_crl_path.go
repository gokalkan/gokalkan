package gokalkan

// GetCRLCachePath возвращает путь к файлу из кэша CRL.
func (cli *Client) GetCRLCachePath(keyUsage KeyUsage) string {
	switch keyUsage {
	case KeyUsageSign:
		return cli.o.CRLCachePathGOST
	case KeyUsageAuth:
		return cli.o.CRLCachePathRSA
	case KeyUsageUnknown:
		return ""
	default:
		return ""
	}
}
