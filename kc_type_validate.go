package gokalkan

// KCValidateType - это тип валидации сертификата
type KCValidateType int

// Константы, определяющие тип валидации
const (
	KCValidateTypeNothing KCValidateType = 1025 // Не делать проверок
	KCValidateTypeCRL     KCValidateType = 1026 // Проверка сертификата по списку отозванных сертификатов
	KCValidateTypeOCSP    KCValidateType = 1028 // Проверка сертификата посредством сервиса OCSP
	// KCFlagGetOCSPResponse KCValidateType = 524288 // Получить ответ от OCSP-сервиса
)
