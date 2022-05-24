package gokalkan

type KCCertType int // Принадлежность сертификата

// Константы, определяющие принадлежность сертификата
const (
	KCCertTypeCA           KCCertType = 513 // Корневой сертификат УЦ
	KCCertTypeIntermediate KCCertType = 514 // Сертификат промежуточного УЦ
	KCCertTypeUser         KCCertType = 516 // Сертификат пользователя
)
