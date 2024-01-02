package gokalkan

import (
	"crypto/x509"
	"encoding/pem"
	"reflect"
	"strings"
	"time"

	"github.com/gokalkan/gokalkan/ckalkan"
	"github.com/gokalkan/gokalkan/types"
)

var certPropMap = map[string]ckalkan.CertProp{
	"Subject":           ckalkan.CertPropSubjectDN,
	"SerialNumber":      ckalkan.CertPropCertCN,
	"ValidFrom":         ckalkan.CertPropNotBefore,
	"ValidUntil":        ckalkan.CertPropNotAfter,
	"Issuer":            ckalkan.CertPropIssuerDN,
	"Policy":            ckalkan.CertPropPoliciesID,
	"KeyUsage":          ckalkan.CertPropKeyUsage,
	"ExtKeyUsage":       ckalkan.CertPropExtKeyUsage,
	"AuthKeyID":         ckalkan.CertPropAuthKeyID,
	"SubjKeyID":         ckalkan.CertPropSubjKeyID,
	"AlgorithmSignCert": ckalkan.CertPropSignatureAlg,
	"PublicKey":         ckalkan.CertPropPubKey,
	"OcspUrl":           ckalkan.CertPropOCSP,
	// "CrlUrl":            ckalkan.CertPropGetCRL,
	// "DeltaCrlUrl":       ckalkan.CertPropGetDeltaCRL,
}

// GetTimeFromSig получает время подписания из CMS в кодировке DER.
// Если вы хотите пeредать данные подписи в base64 формате, то установите флаг base64 = true
func (cli *Client) GetTimeFromSig(cmsDer []byte) (time.Time, error) {
	var flags ckalkan.Flag

	flags = ckalkan.FlagInDER

	return cli.kc.GetTimeFromSig(string(cmsDer), flags, 0)
}

// GetSigAlgFromXML обеспечивает получение алгоритма подписи из XML.
func (cli *Client) GetSigAlgFromXML(xml string) (string, error) {
	return cli.kc.GetSigAlgFromXML(xml)
}

// X509CertificateGetInfo Обеспечивает получение значений полей/расширений из сертификата в виде структуры *types.CertificateInfo.
func (cli *Client) X509CertificateGetInfo(input *x509.Certificate) (*types.CertificateInfo, error) {

	resultCertificateInfo := &types.CertificateInfo{}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: input.Raw})
	certPEMString := string(certPEM)

	value := reflect.ValueOf(resultCertificateInfo).Elem()
	for field, certProp := range certPropMap {
		result, err := cli.kc.X509CertificateGetInfo(certPEMString, certProp)
		if err != nil {
			return nil, err
		}

		switch field {
		case "ValidFrom", "ValidUntil":
			t, err := time.Parse("02.01.2006 15:04:05 MST", strings.Split(result, "=")[1])
			if err != nil {
				return nil, err
			}
			value.FieldByName(field).Set(reflect.ValueOf(t))

		default:
			value.FieldByName(field).SetString(result)
		}
	}

	resultCertificateInfo.Policies = strings.Split(strings.Split(resultCertificateInfo.Policy, "=")[1], ";")
	resultCertificateInfo.KeyUsages = strings.Split(strings.Split(resultCertificateInfo.KeyUsage, "=")[1], ";")
	resultCertificateInfo.ExtKeyUsages = strings.Split(strings.Split(resultCertificateInfo.ExtKeyUsage, "=")[1], ";")

	return resultCertificateInfo, nil
}
