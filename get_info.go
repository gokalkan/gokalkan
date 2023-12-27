package gokalkan

import (
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// GetTimeFromSig получает время подписания из CMS в кодировке DER.
// Если вы хотите пeредать данные подписи в base64 формате, то установите флаг base64 = true
func (cli *Client) GetTimeFromSig(cmsDer []byte) (time.Time, error) {
	var flags ckalkan.Flag

	flags = ckalkan.FlagInDER

	return cli.kc.GetTimeFromSig(string(cmsDer), flags, 0)
}

// X509CertificateGetInfo Обеспечивает получение значений полей/расширений из сертификата
// Сертификат должен быть предварительно загружен с помощью одной из функций: LoadKeyStore(), X509LoadCertificateFromFile().
// Укажите необходимые строки полей в срезе fields из мапы ckalkan.CertPropMap
func (cli *Client) X509CertificateGetInfo(inCert string, fields []string) (string, error) {
	var res string
	for _, field := range fields {
		prop := ckalkan.CertPropMap[field]
		result, err := cli.kc.X509CertificateGetInfo(inCert, prop)
		if err != nil {
			return "", err
		}
		res += result + "\n"
	}
	return res[:len(res)-1], nil
}

// GetSigAlgFromXML обеспечивает получение алгоритма подписи из XML.
func (cli *Client) GetSigAlgFromXML(xml string) (string, error) {
	return cli.kc.GetSigAlgFromXML(xml)
}
