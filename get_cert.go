package gokalkan

import (
	"encoding/base64"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// GetCertFromCMS обеспечивает получение сертификата из CMS.
func (cli *Client) GetCertFromCMS(cms []byte, signID int) (string, error) {
	cmsB64 := base64.StdEncoding.EncodeToString(cms)
	return cli.kc.GetCertFromCMS(cmsB64, signID, ckalkan.FlagInBase64)
}

func (cli *Client) GetCertFromXML(xml []byte, signID int) (string, error) {
	return cli.kc.GetCertFromXML(string(xml), signID)
}
