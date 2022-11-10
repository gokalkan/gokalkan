package ckalkan

// CertType определяет принадлежность сертификата
type CertType int

// Константы, определяющие принадлежность сертификата
const (
	CertTypeCA           CertType = 513 // Корневой сертификат УЦ
	CertTypeIntermediate CertType = 514 // Сертификат промежуточного УЦ
	CertTypeUser         CertType = 516 // Сертификат пользователя
)
