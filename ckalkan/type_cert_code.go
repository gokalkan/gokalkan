package ckalkan

// CertCodeType определяет тип кодировки
type CertCodeType int

// Константы, определяющие тип кодировки
const (
	CertCodeTypeDER    CertCodeType = 257 // Кодировка DER
	CertCodeTypePEM    CertCodeType = 258 // Кодировка PEM
	CertCodeTypeBase64 CertCodeType = 260 // Кодировка Base64
)
