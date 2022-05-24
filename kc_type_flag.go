package gokalkan

type KCFlag int

const (
	KCFlagSignDraft       KCFlag = 1      // Сырая подпись (draft sign)
	KCFlagSignCMS         KCFlag = 2      // Подпись в формате CMS
	KCFlagInPEM           KCFlag = 4      // Входные данные в формате PEM
	KCFlagInDER           KCFlag = 8      // Входные данные в кодировке DER
	KCFlagInBase64        KCFlag = 16     // Входные данные в кодировке BASE64
	KCFlagIn2Base64       KCFlag = 32     // Дополнительные входные данные в кодировке BASE64
	KCFlagDetachedData    KCFlag = 64     // Отсоединенная подпись
	KCFlagWithCert        KCFlag = 128    // Вложить сертификат в подпись
	KCFlagWithTimestamp   KCFlag = 256    // Добавить в подпись метку времени (не используется в текущей версии)
	KCFlagOutPEM          KCFlag = 512    // Выходные данные в формате PEM
	KCFlagOutDER          KCFlag = 1024   // Выходные данные в кодировке DER
	KCFlagOutBase64       KCFlag = 2048   // Выходные данные в кодировке BASE64
	KCFlagProxyOff        KCFlag = 4096   // Отключить использование прокси-сервера и стереть настройки.
	KCFlagProxyOn         KCFlag = 8192   //  Включить и установить настройки прокси-сервера (адрес и порт)
	KCFlagProxyAuth       KCFlag = 16384  // Прокси-сервер требует авторизацию (логин/пароль)
	KCFlagInFile          KCFlag = 32768  // Использовать, если параметр inData/outData содержит абсолютный путь к файлу
	KCFlagNoCheckCertTime KCFlag = 65536  // Не проверять срок действия сертификата при построении цепочки до корневого (для проверки старых подписей с просроченным сертификатом)
	KCFlagHashSHA256      KCFlag = 131072 // Алгоритм хеширования sha256
	KCFlagHashGOST95      KCFlag = 262144 // Алгоритм хеширования Gost34311_95
)
