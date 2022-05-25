package gokalkan

// KCCertType определяет принадлежность сертификата
type KCCertType int

// Константы, определяющие принадлежность сертификата
const (
	KCCertTypeCA           KCCertType = 513 // Корневой сертификат УЦ
	KCCertTypeIntermediate KCCertType = 514 // Сертификат промежуточного УЦ
	KCCertTypeUser         KCCertType = 516 // Сертификат пользователя
)
