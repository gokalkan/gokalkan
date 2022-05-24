package gokalkan

import (
	"strings"
	"time"
)

type X509RawInfo struct {
	NotBefore    string         `json:"not_before"`
	NotAfter     string         `json:"not_after"`
	KeyUsage     string         `json:"key_usage"`
	ExtKeyUsage  string         `json:"ext_key_usage"`
	AuthKeyID    string         `json:"auth_key_id"`
	SubjKeyID    string         `json:"subj_key_id"`
	CertCN       string         `json:"cert_cn"`
	SignatureAlg string         `json:"signature_alg"`
	PublicKey    string         `json:"public_key"`
	PoliciesID   string         `json:"policies_id"`
	OCSP         string         `json:"ocsp"`
	GetCRL       string         `json:"get_crl"`
	GetDeltaCRL  string         `json:"get_delta_crl"`
	Issuer       X509RawIssuer  `json:"issuer"`
	Subject      X509RawSubject `json:"subject"`
}

type X509RawIssuer struct {
	CountryName  string `json:"country_name"`
	SOPN         string `json:"sopn"`
	LocalityName string `json:"locality_name"`
	OrgName      string `json:"org_name"`
	OrgUnitName  string `json:"org_unit_name"`
	CommonName   string `json:"common_name"`
	DN           string `json:"dn"`
}

type X509RawSubject struct {
	CountryName  string `json:"country_name"`
	SOPN         string `json:"sopn"`
	LocalityName string `json:"locality_name"`
	CommonName   string `json:"common_name"`
	GivenName    string `json:"given_name"`
	Surname      string `json:"surname"`
	SerialNumber string `json:"serial_number"`
	Email        string `json:"email"`
	OrgName      string `json:"org_name"`
	OrgUnitName  string `json:"org_unit_name"`
	Bc           string `json:"bc"`
	Dc           string `json:"dc"`
	DN           string `json:"dn"`
}

func (c *X509RawInfo) GetX509() (result X509, err error) {
	na, err := c.GetNotAfter()
	if err != nil {
		return result, err
	}

	nb, err := c.GetNotBefore()
	if err != nil {
		return result, err
	}

	result = X509{
		KeyUsage:  c.GetKeyUsage(),
		KeyUser:   c.GetKeyUser(),
		NotBefore: nb,
		NotAfter:  na,
	}

	return result, nil
}

func (c *X509RawInfo) GetNotBefore() (result time.Time, err error) {
	t := c.NotBefore
	loc := time.Local

	if strings.HasSuffix(t, " ALMT") {
		t = strings.TrimSuffix(t, " ALMT")

		loc, err = time.LoadLocation("Asia/Almaty")
		if err != nil {
			return result, err
		}
	}

	return time.ParseInLocation("02.01.2006 15:04:05", t, loc)
}

func (c *X509RawInfo) GetNotAfter() (result time.Time, err error) {
	t := c.NotAfter
	loc := time.Local

	if strings.HasSuffix(t, " ALMT") {
		t = strings.TrimSuffix(t, " ALMT")

		loc, err = time.LoadLocation("Asia/Almaty")
		if err != nil {
			return result, err
		}
	}

	return time.ParseInLocation("02.01.2006 15:04:05", t, loc)
}

func (c *X509RawInfo) GetKeyUsage() KeyUsage {
	if c == nil {
		return KeyUsageUnknown
	}

	return parseKeyUsage(c.KeyUsage)
}

func (c *X509RawInfo) GetKeyUser() []KeyUser {
	if c == nil {
		return nil
	}

	if c.ExtKeyUsage == "" {
		return nil
	}

	result := make([]KeyUser, 0)

	s := c.ExtKeyUsage

	for k, v := range keyUserMap {
		if strings.Contains(s, "("+v+")") {
			result = append(result, k)
		}
	}

	return result
}
