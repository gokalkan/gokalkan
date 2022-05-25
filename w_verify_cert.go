package gokalkan

import "context"

// VerifyCert - осуществляет проверку сертификата:
// - проверка срока действия,
// - построение цепочки сертификатов,
// - проверка отозванности по OCSP или CRL.
//
//  Если validateType:
//  - KCValidateTypeCRL - в параметр path необходимо указывать путь к файлу crl.
//		Например:
//  	VerifyCert(gostCert, KCValidateTypeCRL, "/tmp/nca_gost.crl")
//
//  - KCValidateTypeOCSP - в параметр path необходимо указывать url OCSP.
// 	  По умолчанию передается url http://ocsp.pki.gov.kz.
//		Например:
//		VerifyCert(gostCert, KCValidateTypeOCSP)
//		VerifyCert(gostCert, KCValidateTypeOCSP, "http://ocsp.pki.gov.kz")
//
//  - KCValidateTypeNothing - не производятся проверки по CRL или OCSP. Параметр path игнорируется.
//  	Например:
//  	VerifyCert(gostCert, KCValidateTypeNothing, "")
func (cli *Client) VerifyCert(cert string, validateType KCValidateType, path ...string) (result string, err error) {
	validatePath := ""

	if len(path) == 0 {
		switch validateType {
		case KCValidateTypeCRL:
			ku, err := cli.GetCertKeyUsage(cert)
			if err != nil {
				return result, err
			}

			err = cli.LoadCRLCache(context.Background())
			if err != nil {
				return result, err
			}

			validatePath = cli.GetCRLCachePath(ku)
		case KCValidateTypeOCSP:
			validatePath = cli.o.OCSP
		case KCValidateTypeNothing:
			validatePath = ""
		default:
		}
	} else {
		switch validateType {
		case KCValidateTypeCRL:
			validatePath = path[0]
		case KCValidateTypeOCSP:
			validatePath = path[0]
		case KCValidateTypeNothing:
			validatePath = ""
		default:
		}
	}

	return cli.kc.KCX509ValidateCertificate(cert, validateType, validatePath)
}
