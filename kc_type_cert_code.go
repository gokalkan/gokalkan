package gokalkan

type KCCertCodeType int // Тип кодировки

// Константы, определяющие тип кодировки
const (
	KCCertCodeTypeDER    KCCertCodeType = 257 // Кодировка DER
	KCCertCodeTypePEM    KCCertCodeType = 258 // Кодировка PEM
	KCCertCodeTypeBase64 KCCertCodeType = 260 // Кодировка Base64
)
