package ckalkan

// Flag представляет собой флаги для KalcanCrypt
type Flag int

const (
	FlagSignDraft       Flag = 1      // Сырая подпись (draft sign)
	FlagSignCMS         Flag = 2      // Подпись в формате CMS
	FlagInPEM           Flag = 4      // Входные данные в формате PEM
	FlagInDER           Flag = 8      // Входные данные в кодировке DER
	FlagInBase64        Flag = 16     // Входные данные в кодировке BASE64
	FlagIn2Base64       Flag = 32     // Дополнительные входные данные в кодировке BASE64
	FlagDetachedData    Flag = 64     // Отсоединенная подпись
	FlagWithCert        Flag = 128    // Вложить сертификат в подпись
	FlagWithTimestamp   Flag = 256    // Добавить в подпись метку времени
	FlagOutPEM          Flag = 512    // Выходные данные в формате PEM
	FlagOutDER          Flag = 1024   // Выходные данные в кодировке DER
	FlagOutBase64       Flag = 2048   // Выходные данные в кодировке BASE64
	FlagProxyOff        Flag = 4096   // Отключить использование прокси-сервера и стереть настройки.
	FlagProxyOn         Flag = 8192   // Включить и установить настройки прокси-сервера (адрес и порт)
	FlagProxyAuth       Flag = 16384  // Прокси-сервер требует авторизацию (логин/пароль)
	FlagInFile          Flag = 32768  // Использовать, если параметр inData/outData содержит абсолютный путь к файлу
	FlagNoCheckCertTime Flag = 65536  // Не проверять срок действия сертификата при построении цепочки до корневого (для проверки старых подписей с просроченным сертификатом)
	FlagHashSHA256      Flag = 131072 // Алгоритм хеширования sha256
	FlagHashGOST95      Flag = 262144 // Алгоритм хеширования Gost34311_95
)
