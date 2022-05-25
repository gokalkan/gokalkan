package gokalkan

import "strings"

// KeyUsage - это назначение ключа. Может быть одним из значении:
//
// - Пустая строка - неопределенное значение.
//
// - SIGN - Ключ предназначен для подписи документов.
//
// - AUTH - Ключ предназначен для аутентификации.
type KeyUsage string

func (k KeyUsage) String() string {
	return string(k)
}

const (
	KeyUsageUnknown KeyUsage = ""     // Это значение устанавливается, если не удалось определит параметр "Использования ключа"
	KeyUsageSign    KeyUsage = "SIGN" // Ключ предназначен для подписи документов
	KeyUsageAuth    KeyUsage = "AUTH" // Ключ предназначен для аутентификации (обычно такие ключи начинаются с AUTH_RSA_256...)
)

func parseKeyUsage(s string) KeyUsage {
	ku := strings.Split(strings.TrimSpace(s), " ")

	digitalSignature := false
	nonRepudiation := false
	keyEncipherment := false

	for i := range ku {
		if strings.EqualFold(ku[i], "digitalSignature") {
			digitalSignature = true
		}

		if strings.EqualFold(ku[i], "nonRepudiation") {
			nonRepudiation = true
		}

		if strings.EqualFold(ku[i], "keyEncipherment") {
			keyEncipherment = true
		}
	}

	if digitalSignature && nonRepudiation {
		return KeyUsageSign
	}

	if digitalSignature && keyEncipherment {
		return KeyUsageAuth
	}

	return KeyUsageUnknown
}
