package gokalkan

import "time"

const (
	CertTypeIndividual CertType = iota
	CertTypeOrganization
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
	CertType        int
	CertSubjectRole int
	CertSubject     struct {
		CommonName string
		LastName   string
		Country    string
		IIN        string
		DN         string
	}
	CertOrganization struct {
		Name         string
		BIN          string
		SubjectRole  CertSubjectRole
		SubjectEmail string
	}
	CertIssuer struct {
		CommonName string
		Country    string
		DN         string
	}
)
