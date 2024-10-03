package gokalkan

import "time"

const (
	CertTypeIndividual   = "individual"
	CertTypeOrganization = "organization"
)

const (
	CertSubjectRoleUndefined CertSubjectRole = iota
	CertSubjectRoleCEO
	CertSubjectRoleSign
	CertSubjectRoleSignFinance
	CertSubjectRoleHR
	CertSubjectRoleEmployee
)

type (
	Summary struct {
		Type         CertType
		Subject      CertSubject
		Organization CertOrganization
		Issuer       CertIssuer
		PublicKey    string
		SerialNumber string
		NotAfter     time.Time
		NotBefore    time.Time
	}
	CertType        string
	CertSubjectRole int
	CertSubject     struct {
		CommonName string
		LastName   string
		Country    string
		IIN        string
		DN         string
	}
	CertOrganization struct {
		Name        string
		BIN         string
		SubjectRole CertSubjectRole
	}
	CertIssuer struct {
		CommonName string
		Country    string
		DN         string
	}
)
