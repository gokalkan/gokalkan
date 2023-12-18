package gokalkan

import (
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
)

// GetTimeFromSig получает время подписания из CMS
// Если вы хотите пeредать данные подписи в base64 формате, то установите флаг base64 = true
func (cli *Client) GetTimeFromSig(data string, base64 bool) (time.Time, error) {
	var flag ckalkan.Flag

	flag = ckalkan.FlagInDER

	if base64 {
		flag = ckalkan.FlagInBase64
	}

	return cli.kc.GetTimeFromSig(data, flag, 0)
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
