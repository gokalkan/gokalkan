package ckalkan

// ValidateType - это тип валидации сертификата
type ValidateType int

// Константы, определяющие тип валидации
const (
	ValidateTypeNothing     ValidateType = 1025   // Не делать проверок
	ValidateTypeCRL         ValidateType = 1026   // Проверка сертификата по списку отозванных сертификатов
	ValidateTypeOCSP        ValidateType = 1028   // Проверка сертификата посредством сервиса OCSP
	ValidateGetOCSPResponse ValidateType = 524288 // Получить ответ от OCSP-сервиса
)
