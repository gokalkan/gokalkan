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

func (cli *Client) getCertProp(certPEM string, prop KCCertProp) (string, error) {
	v, err := cli.kc.KCX509CertificateGetInfo(certPEM, prop)
	if err != nil {
		return "", err
	}

	switch prop {
	case KCCertPropIssuerCountryName:
		return strings.TrimPrefix(v, "C="), nil
	case KCCertPropIssuerSOPN:
		return v, nil
	case KCCertPropIssuerLocalityName:
		return v, nil
	case KCCertPropIssuerOrgName:
		return strings.TrimPrefix(v, "O="), nil
	case KCCertPropIssuerOrgUnitName:
		return strings.TrimPrefix(v, "OU="), nil
	case KCCertPropIssuerCommonName:
		return strings.TrimPrefix(v, "CN="), nil
	case KCCertPropSubjectCountryName:
		return strings.TrimPrefix(v, "C="), nil
	case KCCertPropSubjectSOPN:
		return v, nil
	case KCCertPropSubjectLocalityName:
		return v, nil
	case KCCertPropSubjectCommonName:
		return strings.TrimPrefix(v, "CN="), nil
	case KCCertPropSubjectGivenName:
		return strings.TrimPrefix(v, "GN="), nil
	case KCCertPropSubjectSurname:
		return strings.TrimPrefix(v, "SN="), nil
	case KCCertPropSubjectSerialNumber:
		return strings.TrimPrefix(v, "serialNumber="), nil
	case KCCertPropSubjectEmail:
		return strings.TrimPrefix(v, "emailAddress="), nil
	case KCCertPropSubjectOrgName:
		return strings.TrimPrefix(v, "O="), nil
	case KCCertPropSubjectOrgUnitName:
		return strings.TrimPrefix(v, "OU="), nil
	case KCCertPropSubjectBc:
		return v, nil
	case KCCertPropSubjectDc:
		return v, nil
	case KCCertPropNotBefore:
		return strings.TrimPrefix(v, "notBefore="), nil
	case KCCertPropNotAfter:
		return strings.TrimPrefix(v, "notAfter="), nil
	case KCCertPropKeyUsage:
		return strings.TrimPrefix(v, "keyUsage="), nil
	case KCCertPropExtKeyUsage:
		return strings.TrimPrefix(v, "extendedKeyUsage="), nil
	case KCCertPropAuthKeyID:
		return strings.TrimPrefix(v, "authorityKeyIdentifier="), nil
	case KCCertPropSubjKeyID:
		return strings.TrimPrefix(v, "subjectKeyIdentifier="), nil
	case KCCertPropCertCN:
		return strings.TrimPrefix(v, "certificateSerialNumber="), nil
	case KCCertPropIssuerDN:
		return v, nil
	case KCCertPropSubjectDN:
		return v, nil
	case KCCertPropSignatureAlg:
		return strings.TrimPrefix(v, "signatureAlgorithm="), nil
	case KCCertPropPubKey:
		return v, nil
	case KCCertPropPoliciesID:
		return strings.TrimPrefix(v, "certificatePolicies="), nil
	case KCCertPropOCSP:
		return strings.TrimPrefix(v, "OCSP="), nil
	case KCCertPropGetCRL:
		return strings.TrimPrefix(v, "crlDistributionPoints="), nil
	case KCCertPropGetDeltaCRL:
		return strings.TrimPrefix(v, "freshestCRL="), nil
	default:
	}
	return "", nil
}

//nolint:gocyclo
func (cli *Client) setCertInfo(certPEM string, prop KCCertProp, crt *X509RawInfo) error {
	v, err := cli.getCertProp(certPEM, prop)
	if err != nil {
		return err
	}

	switch prop {
	case KCCertPropIssuerCountryName:
		crt.Issuer.CountryName = v
	case KCCertPropIssuerSOPN:
		crt.Issuer.SOPN = v
	case KCCertPropIssuerLocalityName:
		crt.Issuer.LocalityName = v
	case KCCertPropIssuerOrgName:
		crt.Issuer.OrgName = v
	case KCCertPropIssuerOrgUnitName:
		crt.Issuer.OrgUnitName = v
	case KCCertPropIssuerCommonName:
		crt.Issuer.CommonName = v
	case KCCertPropSubjectCountryName:
		crt.Subject.CountryName = v
	case KCCertPropSubjectSOPN:
		crt.Subject.SOPN = v
	case KCCertPropSubjectLocalityName:
		crt.Subject.LocalityName = v
	case KCCertPropSubjectCommonName:
		crt.Subject.CommonName = v
	case KCCertPropSubjectGivenName:
		crt.Subject.GivenName = v
	case KCCertPropSubjectSurname:
		crt.Subject.Surname = v
	case KCCertPropSubjectSerialNumber:
		crt.Subject.SerialNumber = v
	case KCCertPropSubjectEmail:
		crt.Subject.Email = v
	case KCCertPropSubjectOrgName:
		crt.Subject.OrgName = v
	case KCCertPropSubjectOrgUnitName:
		crt.Subject.OrgUnitName = v
	case KCCertPropSubjectBc:
		crt.Subject.OrgUnitName = v
	case KCCertPropSubjectDc:
		crt.Subject.Dc = v
	case KCCertPropNotBefore:
		crt.NotBefore = v
	case KCCertPropNotAfter:
		crt.NotAfter = v
	case KCCertPropKeyUsage:
		crt.KeyUsage = v
	case KCCertPropExtKeyUsage:
		crt.ExtKeyUsage = v
	case KCCertPropAuthKeyID:
		crt.AuthKeyID = v
	case KCCertPropSubjKeyID:
		crt.SubjKeyID = v
	case KCCertPropCertCN:
		crt.CertCN = v
	case KCCertPropIssuerDN:
		crt.Issuer.DN = v
	case KCCertPropSubjectDN:
		crt.Subject.DN = v
	case KCCertPropSignatureAlg:
		crt.SignatureAlg = v
	case KCCertPropPubKey:
		crt.PublicKey = v
	case KCCertPropPoliciesID:
		crt.PoliciesID = v
	case KCCertPropOCSP:
		crt.OCSP = v
	case KCCertPropGetCRL:
		crt.GetCRL = v
	case KCCertPropGetDeltaCRL:
		crt.GetDeltaCRL = v
	default:
	}

	return nil
}
