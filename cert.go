package gokalkan

import (
	"github.com/gokalkan/gokalkan/ckalkan"
	"strings"
	"time"
)

// see: https://adilet.zan.kz/rus/docs/V2000021440
const (
	oidSubjectIndividual      = "1.2.398.3.3.4.1.1"
	oidSubjectRoleCEO         = "1.2.398.3.3.4.1.2.1"
	oidSubjectRoleSign        = "1.2.398.3.3.4.1.2.2"
	oidSubjectRoleSignFinance = "1.2.398.3.3.4.1.2.3"
	oidSubjectRoleHR          = "1.2.398.3.3.4.1.2.4"
	oidSubjectRoleEmployee    = "1.2.398.3.3.4.1.2.5"
)

const timeLayout = "02.01.2006 15:04:05 -07:00"

func (cli *Client) X509CertificateGetInfo(cert string, prop ckalkan.CertProp) (string, error) {
	return cli.kc.X509CertificateGetInfo(cert, prop)
}

// X509CertificateGetSummary возвращает информацию о сертификате.
// Используйте только после Verify, если используется для проверки подписей
func (cli *Client) X509CertificateGetSummary(cert string) (*Summary, error) {
	var (
		summary = Summary{}
		err     error
	)

	if summary.Subject.CommonName, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropSubjectCommonName,
	); err != nil {
		return nil, err
	}

	if summary.Subject.LastName, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropSubjectGivenName,
	); err != nil {
		return nil, err
	}

	if summary.Subject.Country, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropSubjectCountryName,
	); err != nil {
		return nil, err
	}

	if summary.Subject.IIN, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropSubjectSerialNumber,
	); err != nil {
		return nil, err
	}

	summary.Subject.IIN = cleanupValue(summary.Subject.IIN, "IIN")

	if summary.Subject.DN, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropSubjectDN,
	); err != nil {
		return nil, err
	}

	ekum := make(map[string]bool)
	extKeyUsage := ""

	if extKeyUsage, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropExtKeyUsage,
	); err != nil {
		return nil, err
	}

	entries := strings.Split(extKeyUsage, ";")

	for _, entry := range entries {
		entry = strings.TrimSpace(entry)
		if entry == "" {
			continue
		}

		if value, found := extractValueInBrackets(entry); found {
			ekum[value] = true
		}
	}

	if _, found := ekum[oidSubjectIndividual]; found {
		summary.Type = CertTypeIndividual
	} else {
		summary.Type = CertTypeOrganization
	}

	if summary.Type == CertTypeOrganization {
		if summary.Organization.Name, err = cli.kc.X509CertificateGetInfo(
			cert,
			ckalkan.CertPropSubjectOrgName,
		); err != nil {
			return nil, err
		}

		if summary.Organization.BIN, err = cli.kc.X509CertificateGetInfo(
			cert,
			ckalkan.CertPropSubjectSerialNumber,
		); err != nil {
			return nil, err
		}

		if parts := strings.Split(summary.Organization.BIN, "BIN"); len(parts) == 2 {
			summary.Organization.BIN = parts[1]
		}

		if _, exists := ekum[oidSubjectRoleCEO]; exists {
			summary.Organization.SubjectRole = CertSubjectRoleCEO
		} else if _, exists := ekum[oidSubjectRoleSign]; exists {
			summary.Organization.SubjectRole = CertSubjectRoleSign
		} else if _, exists := ekum[oidSubjectRoleSignFinance]; exists {
			summary.Organization.SubjectRole = CertSubjectRoleSignFinance
		} else if _, exists := ekum[oidSubjectRoleHR]; exists {
			summary.Organization.SubjectRole = CertSubjectRoleHR
		} else if _, exists := ekum[oidSubjectRoleEmployee]; exists {
			summary.Organization.SubjectRole = CertSubjectRoleEmployee
		}
	}

	if summary.Issuer.CommonName, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropIssuerCommonName,
	); err != nil {
		return nil, err
	}

	if summary.Issuer.Country, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropIssuerCountryName,
	); err != nil {
		return nil, err
	}

	if summary.Issuer.DN, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropIssuerDN,
	); err != nil {
		return nil, err
	}

	if summary.PublicKey, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropPubKey,
	); err != nil {
		return nil, err
	}

	if summary.SerialNumber, err = cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropCertCN,
	); err != nil {
		return nil, err
	}

	if parts := strings.Split(summary.SerialNumber, "="); len(parts) == 2 {
		summary.SerialNumber = parts[1]
	}

	if notAfter, err := cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropNotAfter,
	); err != nil {
		return nil, err
	} else if part := strings.Split(notAfter, "="); len(part) == 2 {
		if summary.NotAfter, err = time.Parse(timeLayout, part[1]); err != nil {
			return nil, err
		}
	}

	if notBefore, err := cli.kc.X509CertificateGetInfo(
		cert,
		ckalkan.CertPropNotBefore,
	); err != nil {
		return nil, err
	} else if part := strings.Split(notBefore, "="); len(part) == 2 {
		if summary.NotBefore, err = time.Parse(timeLayout, part[1]); err != nil {
			return nil, err
		}
	}

	return &summary, nil
}

func extractValueInBrackets(entry string) (string, bool) {
	start := strings.Index(entry, "(")
	end := strings.LastIndex(entry, ")")

	if start != -1 && end != -1 && start < end {
		return strings.TrimSpace(entry[start+1 : end]), true
	}

	return "", false
}

func cleanupValue(value string, s string) string {
	if parts := strings.Split(value, s); len(parts) == 2 {
		return parts[1]
	}

	return value
}
