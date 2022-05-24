package gokalkan

import "strings"

// GetCertInfo возвращает информацию о сертификате.
func (cli *Client) GetCertInfo(certPEM string) (result *X509RawInfo, err error) {
	result = &X509RawInfo{}

	for i := range allProps {
		err = cli.setCertInfo(certPEM, allProps[i], result)
		if err != nil {
			code, ok := GetErrorCode(err)
			if ok && code == ErrorCodeGetCertPropErr {
				continue
			}

			return result, err
		}
	}

	return result, nil
}

//nolint:gocyclo
func (cli *Client) setCertInfo(certPEM string, prop KCCertProp, crt *X509RawInfo) error {
	v, err := cli.kc.KCX509CertificateGetInfo(certPEM, prop)
	if err != nil {
		return err
	}

	switch prop {
	case KCCertPropIssuerCountryName:
		crt.Issuer.CountryName = strings.TrimPrefix(v, "C=")
	case KCCertPropIssuerSOPN:
		crt.Issuer.SOPN = v
	case KCCertPropIssuerLocalityName:
		crt.Issuer.LocalityName = v
	case KCCertPropIssuerOrgName:
		crt.Issuer.OrgName = strings.TrimPrefix(v, "O=")
	case KCCertPropIssuerOrgUnitName:
		crt.Issuer.OrgUnitName = strings.TrimPrefix(v, "OU=")
	case KCCertPropIssuerCommonName:
		crt.Issuer.CommonName = strings.TrimPrefix(v, "CN=")
	case KCCertPropSubjectCountryName:
		crt.Subject.CountryName = strings.TrimPrefix(v, "C=")
	case KCCertPropSubjectSOPN:
		crt.Subject.SOPN = v
	case KCCertPropSubjectLocalityName:
		crt.Subject.LocalityName = v
	case KCCertPropSubjectCommonName:
		crt.Subject.CommonName = strings.TrimPrefix(v, "CN=")
	case KCCertPropSubjectGivenName:
		crt.Subject.GivenName = strings.TrimPrefix(v, "GN=")
	case KCCertPropSubjectSurname:
		crt.Subject.Surname = strings.TrimPrefix(v, "SN=")
	case KCCertPropSubjectSerialNumber:
		crt.Subject.SerialNumber = strings.TrimPrefix(v, "serialNumber=")
	case KCCertPropSubjectEmail:
		crt.Subject.Email = strings.TrimPrefix(v, "emailAddress=")
	case KCCertPropSubjectOrgName:
		crt.Subject.OrgName = strings.TrimPrefix(v, "O=")
	case KCCertPropSubjectOrgUnitName:
		crt.Subject.OrgUnitName = strings.TrimPrefix(v, "OU=")
	case KCCertPropSubjectBc:
		crt.Subject.OrgUnitName = v
	case KCCertPropSubjectDc:
		crt.Subject.Dc = v
	case KCCertPropNotBefore:
		crt.NotBefore = strings.TrimPrefix(v, "notBefore=")
	case KCCertPropNotAfter:
		crt.NotAfter = strings.TrimPrefix(v, "notAfter=")
	case KCCertPropKeyUsage:
		crt.KeyUsage = strings.TrimPrefix(v, "keyUsage=")
	case KCCertPropExtKeyUsage:
		crt.ExtKeyUsage = strings.TrimPrefix(v, "extendedKeyUsage=")
	case KCCertPropAuthKeyID:
		crt.AuthKeyID = strings.TrimPrefix(v, "authorityKeyIdentifier=")
	case KCCertPropSubjKeyID:
		crt.SubjKeyID = strings.TrimPrefix(v, "subjectKeyIdentifier=")
	case KCCertPropCertCN:
		crt.CertCN = strings.TrimPrefix(v, "certificateSerialNumber=")
	case KCCertPropIssuerDN:
		crt.Issuer.DN = v
	case KCCertPropSubjectDN:
		crt.Subject.DN = v
	case KCCertPropSignatureAlg:
		crt.SignatureAlg = strings.TrimPrefix(v, "signatureAlgorithm=")
	case KCCertPropPubKey:
		crt.PublicKey = v
	case KCCertPropPoliciesID:
		crt.PoliciesID = strings.TrimPrefix(v, "certificatePolicies=")
	case KCCertPropOCSP:
		crt.OCSP = strings.TrimPrefix(v, "OCSP=")
	case KCCertPropGetCRL:
		crt.GetCRL = strings.TrimPrefix(v, "crlDistributionPoints=")
	case KCCertPropGetDeltaCRL:
		crt.GetDeltaCRL = strings.TrimPrefix(v, "freshestCRL=")
	default:
	}

	return nil
}
