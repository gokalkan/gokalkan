package gokalkan

import "strconv"

type KCErrorCode int64 // Код ошибки времени выполнения

func (t KCErrorCode) String() string {
	v, ok := errorLabels[t]
	if ok {
		return v
	}

	return "Неизвестный код ошибки"
}

const baseHex = 16

func (t KCErrorCode) Hex() string {
	return strconv.FormatInt(int64(t), baseHex)
}

const (
	ErrorCodeOK                           KCErrorCode = 0         // Нет ошибки
	ErrorCodeErrorReadPKCS12              KCErrorCode = 149946370 // Невозможно прочитать файл формата pkcs#12
	ErrorCodeErrorOpenPKCS12              KCErrorCode = 149946371 // Невозможно открыть файл формата pkcs12
	ErrorCodeInvalidPropID                KCErrorCode = 149946372 // Недопустимый идентификатор расширения сертификата
	ErrorCodeBufferTooSmall               KCErrorCode = 149946373 // Размер буфера слишком мал
	ErrorCodeCertParseError               KCErrorCode = 149946374 // Невозможно разобрать (распарсить) сертификат
	ErrorCodeInvalidFlag                  KCErrorCode = 149946375 // Недопустимый флаг
	ErrorCodeOpenFileErr                  KCErrorCode = 149946376 // Невозможно открыть файл
	ErrorCodeInvalidPassword              KCErrorCode = 149946377 // Неправильный пароль
	ErrorCodeMemoryError                  KCErrorCode = 149946381 // Невозможно выделить память
	ErrorCodeCheckChainError              KCErrorCode = 149946382 // Не найден сертификат УЦ или сертификат пользователя при проверки цепочки
	ErrorCodeValidTypeError               KCErrorCode = 149946384 // Недопустимый тип валидации сертификата
	ErrorCodeBadCRLFormat                 KCErrorCode = 149946385 // Некорректный формат CRL
	ErrorCodeLoadCRLError                 KCErrorCode = 149946386 // Невозможно загрузить CRL
	ErrorCodeLoadCRLsError                KCErrorCode = 149946387 // Невозможно загрузить CRL-ы
	ErrorCodeUnknownAlg                   KCErrorCode = 149946389 // Неизвестный алгоритм подписи
	ErrorCodeKeyNotFound                  KCErrorCode = 149946390 // Не найден приватный ключ пользователя
	ErrorCodeSignInitError                KCErrorCode = 149946391 // Невозможно инициализировать менеджера подписи
	ErrorCodeSignError                    KCErrorCode = 149946392 // Не удалось сгенерировать цифровую подпись
	ErrorCodeEncodeError                  KCErrorCode = 149946393 // Ошибка шифрования
	ErrorCodeInvalidFlags                 KCErrorCode = 149946394 // Недопустимые флаги
	ErrorCodeCertNotFound                 KCErrorCode = 149946395 // Не найден сертификат пользователя
	ErrorCodeVerifySignError              KCErrorCode = 149946396 // Ошибка верификации подписи xml
	ErrorCodeBase64DecodeError            KCErrorCode = 149946397 // Ошибка дешифровки из Base 64
	ErrorCodeUnknownCMSFormat             KCErrorCode = 149946398 // Неизвестный формат CMS
	ErrorCodeCACertNotFound               KCErrorCode = 149946400 // Не найден сертификат УЦ
	ErrorCodeXMLSecInitError              KCErrorCode = 149946401 // Ошибка инициализации xmlsec
	ErrorCodeLoadTrustedCertsErr          KCErrorCode = 149946402 // Ошибка загрузки доверенных сертификатов
	ErrorCodeSignInvalid                  KCErrorCode = 149946403 // Недопустимая подпись xml
	ErrorCodeNoSignFound                  KCErrorCode = 149946404 // Не найдена подпись во входных данных
	ErrorCodeDecodeError                  KCErrorCode = 149946405 // Ошибка дешифрования
	ErrorCodeXMLParseError                KCErrorCode = 149946406 // Невозможно разобрать (распарсить) xml
	ErrorCodeXMLAddIDError                KCErrorCode = 149946407 // Не удалось добавить атрибут ID
	ErrorCodeXMLInternalError             KCErrorCode = 149946408 // Ошибка при работе с xml
	ErrorCodeXMLSetSignError              KCErrorCode = 149946409 // Не удалось подписать xml
	ErrorCodeOpenSSLError                 KCErrorCode = 149946410 // Ошибка openssl
	ErrorCodeNoTokenFound                 KCErrorCode = 149946412 // Не найден токен
	ErrorCodeOCSPAddCertErr               KCErrorCode = 149946413 // Не удалось добавить сертификат в ocsp
	ErrorCodeOCSPParseURLErr              KCErrorCode = 149946414 // Не удалось разобрать url
	ErrorCodeOCSPAddHostErr               KCErrorCode = 149946415 // Не удалось добавить хост
	ErrorCodeOCSPReqErr                   KCErrorCode = 149946416 // Не удалось добавить текущее время в запрос
	ErrorCodeOCSPConnectionErr            KCErrorCode = 149946417 // Ошибка подключения к OCSP респондеру
	ErrorCodeVerifyNoData                 KCErrorCode = 149946418 // Нет входных данных для верификации
	ErrorCodeIDAttrNotFound               KCErrorCode = 149946419 // Не найден атрибут ID
	ErrorCodeIDRange                      KCErrorCode = 149946420 // Некорректный идентификатор
	ErrorCodeReaderNotFound               KCErrorCode = 149946423 // Не найден ридер
	ErrorCodeGetCertPropErr               KCErrorCode = 149946424 // Не удалось получить значение атрибута
	ErrorCodeSignFormat                   KCErrorCode = 149946425 // Неизвестный формат подписи
	ErrorCodeInDataFormat                 KCErrorCode = 149946426 // Неизвестный формат входных данных
	ErrorCodeOutDataFormat                KCErrorCode = 149946427 // Неизвестный формат выходных данных
	ErrorCodeVerifyInitError              KCErrorCode = 149946428 // Невозможно инициализировать менеджера верификации подписи
	ErrorCodeVerifyError                  KCErrorCode = 149946429 // Не удалось верифицировать цифровую подпись
	ErrorCodeHashError                    KCErrorCode = 149946430 // Не удалось хэшировать данные
	ErrorCodeSignHashError                KCErrorCode = 149946431 // Не удалось подписать хэшированные данные
	ErrorCodeCACertsNotFound              KCErrorCode = 149946432 // Не найден сертификат УЦ в хранилище сертификатов
	ErrorCodeCertTimeInvalid              KCErrorCode = 149946434 // Срок действия сертификата истек либо еще не наступил
	ErrorCodeConvertError                 KCErrorCode = 149946435 // Ошибка записи сертификата в структуру X509
	ErrorCodeTSACreateQuery               KCErrorCode = 149946436 // Ошибка генерации запроса timestamp
	ErrorCodeCreateObj                    KCErrorCode = 149946437 // Ошибка записи OID в ASN1 структуру
	ErrorCodeCreateNoNce                  KCErrorCode = 149946438 // Ошибка генерации уникального числа
	ErrorCodeHTTPError                    KCErrorCode = 149946439 // Ошибка протокола http
	ErrorCodeCADESBESFailed               KCErrorCode = 149946440 // Ошибка проверки расширения CADESBES в CMS
	ErrorCodeCADESTFailed                 KCErrorCode = 149946441 // Ошибка проверки подписи токена TSA
	ErrorCodeNoTSAToken                   KCErrorCode = 149946442 // В подписи не присутствует метка TSA
	ErrorCodeInvalidDigestLen             KCErrorCode = 149946443 // Неправильная длина хэша
	ErrorCodeGenRandError                 KCErrorCode = 149946444 // Ошибка генерации случайного числа
	ErrorCodeSoapNSError                  KCErrorCode = 149946445 // Не найдены заголовки SOAP-сообщений
	ErrorCodeGetPubKey                    KCErrorCode = 149946446 // Ошибка экспорта публичного ключа
	ErrorCodeGetCertInfo                  KCErrorCode = 149946447 // Ошибка получения информации о сертификате
	ErrorCodeFileReadError                KCErrorCode = 149946448 // Ошибка чтения файла
	ErrorCodeCheckError                   KCErrorCode = 149946449 // Хэш не совпадает
	ErrorCodeZipExtractErr                KCErrorCode = 149946450 // Невозможно открыть архив
	ErrorCodeNoManifestFile               KCErrorCode = 149946451 // Не найден MANIFEST
	ErrorCodeVerifyTSHash                 KCErrorCode = 149946452 // не удалось проверить Хэш подписи TS
	ErrorCodeXADESTFailed                 KCErrorCode = 149946453 // XAdES-T: Ошибка проверки подписи
	ErrorCodeOCSPRespStatMalformedRequest KCErrorCode = 149946454 // Неправильный запрос
	ErrorCodeOCSPRespStatInternalError    KCErrorCode = 149946455 // Внутренняя ошибка
	ErrorCodeOCSPRespStatTryLater         KCErrorCode = 149946456 // Попробуйте позже
	ErrorCodeOCSPRespStatSigRequired      KCErrorCode = 149946457 // Должны подписать запрос
	ErrorCodeOCSPRespStatUnauthorized     KCErrorCode = 149946458 // Запрос не авторизован
	ErrorCodeVerifyIssuerSerialV2         KCErrorCode = 149946459 // не удалось проверить IssuerSerialV2 в XAdES
	ErrorCodeOCSPCheckCertFromResp        KCErrorCode = 149946460 // Ошибка проверки сертификата OCSP-респондера
	ErrorCodeCRLExpired                   KCErrorCode = 149946461 // CRL-файл просрочен
	ErrorCodeLibraryNotInitialized        KCErrorCode = 149946625 // Библиотека не инициализирована
	ErrorCodeEngineLoadErr                KCErrorCode = 149946880 // Ошибка подключения (загрузки) модуля (engine)
	ErrorCodeParamError                   KCErrorCode = 149947136 // Некорректные входные данные
	ErrorCodeCertStatusOK                 KCErrorCode = 149947392 // Статус сертификата – валидный. Используется при проверке сертификата по OCSP. (не является ошибкой, делается запись в лог)
	ErrorCodeCertStatusRevoked            KCErrorCode = 149947393 // Статус сертификата – отозван. Используется при проверке сертификата по OCSP.
	ErrorCodeCertStatusUnknown            KCErrorCode = 149947394 // Статус сертификата – неизвестен. Используется при проверке сертификата по OCSP. Например, не удалось установить издателя сертификата.
)

var errorLabels = map[KCErrorCode]string{
	ErrorCodeOK:                           "Нет ошибки",
	ErrorCodeErrorReadPKCS12:              "Невозможно прочитать файл формата pkcs#12",
	ErrorCodeErrorOpenPKCS12:              "Невозможно открыть файл формата pkcs12",
	ErrorCodeInvalidPropID:                "Недопустимый идентификатор расширения сертификата",
	ErrorCodeBufferTooSmall:               "Размер буфера слишком мал",
	ErrorCodeCertParseError:               "Невозможно разобрать (распарсить) сертификат",
	ErrorCodeInvalidFlag:                  "Недопустимый флаг",
	ErrorCodeOpenFileErr:                  "Невозможно открыть файл",
	ErrorCodeInvalidPassword:              "Неправильный пароль",
	ErrorCodeMemoryError:                  "Невозможно выделить память",
	ErrorCodeCheckChainError:              "Не найден сертификат УЦ или сертификат пользователя при проверки цепочки",
	ErrorCodeValidTypeError:               "Недопустимый тип валидации сертификата",
	ErrorCodeBadCRLFormat:                 "Некорректный формат CRL",
	ErrorCodeLoadCRLError:                 "Невозможно загрузить CRL",
	ErrorCodeLoadCRLsError:                "Невозможно загрузить CRL-ы",
	ErrorCodeUnknownAlg:                   "Неизвестный алгоритм подписи",
	ErrorCodeKeyNotFound:                  "Не найден приватный ключ пользователя",
	ErrorCodeSignInitError:                "Невозможно инициализировать менеджера подписи",
	ErrorCodeSignError:                    "Не удалось сгенерировать цифровую подпись",
	ErrorCodeEncodeError:                  "Ошибка шифрования",
	ErrorCodeInvalidFlags:                 "Недопустимые флаги",
	ErrorCodeCertNotFound:                 "Не найден сертификат пользователя",
	ErrorCodeVerifySignError:              "Ошибка верификации подписи xml",
	ErrorCodeBase64DecodeError:            "Ошибка дешифровки из Base 64",
	ErrorCodeUnknownCMSFormat:             "Неизвестный формат CMS",
	ErrorCodeCACertNotFound:               "Не найден сертификат УЦ",
	ErrorCodeXMLSecInitError:              "Ошибка инициализации xmlsec",
	ErrorCodeLoadTrustedCertsErr:          "Ошибка загрузки доверенных сертификатов",
	ErrorCodeSignInvalid:                  "Недопустимая подпись xml",
	ErrorCodeNoSignFound:                  "Не найдена подпись во входных данных",
	ErrorCodeDecodeError:                  "Ошибка дешифрования",
	ErrorCodeXMLParseError:                "Невозможно разобрать (распарсить) xml",
	ErrorCodeXMLAddIDError:                "Не удалось добавить атрибут ID",
	ErrorCodeXMLInternalError:             "Ошибка при работе с xml",
	ErrorCodeXMLSetSignError:              "Не удалось подписать xml",
	ErrorCodeOpenSSLError:                 "Ошибка openssl",
	ErrorCodeNoTokenFound:                 "Не найден токен",
	ErrorCodeOCSPAddCertErr:               "Не удалось добавить сертификат в ocsp",
	ErrorCodeOCSPParseURLErr:              "Не удалось разобрать url",
	ErrorCodeOCSPAddHostErr:               "Не удалось добавить хост",
	ErrorCodeOCSPReqErr:                   "Не удалось добавить текущее время в запрос",
	ErrorCodeOCSPConnectionErr:            "Ошибка подключения к OCSP респондеру",
	ErrorCodeVerifyNoData:                 "Нет входных данных для верификации",
	ErrorCodeIDAttrNotFound:               "Не найден атрибут ID",
	ErrorCodeIDRange:                      "Некорректный идентификатор",
	ErrorCodeReaderNotFound:               "Не найден ридер",
	ErrorCodeGetCertPropErr:               "Не удалось получить значение атрибута",
	ErrorCodeSignFormat:                   "Неизвестный формат подписи",
	ErrorCodeInDataFormat:                 "Неизвестный формат входных данных",
	ErrorCodeOutDataFormat:                "Неизвестный формат выходных данных",
	ErrorCodeVerifyInitError:              "Невозможно инициализировать менеджера верификации подписи",
	ErrorCodeVerifyError:                  "Не удалось верифицировать цифровую подпись",
	ErrorCodeHashError:                    "Не удалось хэшировать данные",
	ErrorCodeSignHashError:                "Не удалось подписать хэшированные данные",
	ErrorCodeCACertsNotFound:              "Не найден сертификат УЦ в хранилище сертификатов",
	ErrorCodeCertTimeInvalid:              "Срок действия сертификата истек либо еще не наступил",
	ErrorCodeConvertError:                 "Ошибка записи сертификата в структуру X509",
	ErrorCodeTSACreateQuery:               "Ошибка генерации запроса timestamp",
	ErrorCodeCreateObj:                    "Ошибка записи OID в ASN1 структуру",
	ErrorCodeCreateNoNce:                  "Ошибка генерации уникального числа",
	ErrorCodeHTTPError:                    "Ошибка протокола http",
	ErrorCodeCADESBESFailed:               "Ошибка проверки расширения CADESBES в CMS",
	ErrorCodeCADESTFailed:                 "Ошибка проверки подписи токена TSA",
	ErrorCodeNoTSAToken:                   "В подписи не присутствует метка TSA",
	ErrorCodeInvalidDigestLen:             "Неправильная длина хэша",
	ErrorCodeGenRandError:                 "Ошибка генерации случайного числа",
	ErrorCodeSoapNSError:                  "Не найдены заголовки SOAP-сообщений",
	ErrorCodeGetPubKey:                    "Ошибка экспорта публичного ключа",
	ErrorCodeGetCertInfo:                  "Ошибка получения информации о сертификате",
	ErrorCodeFileReadError:                "Ошибка чтения файла",
	ErrorCodeCheckError:                   "Хэш не совпадает",
	ErrorCodeZipExtractErr:                "Невозможно открыть архив",
	ErrorCodeNoManifestFile:               "Не найден MANIFEST",
	ErrorCodeVerifyTSHash:                 "не удалось проверить Хэш подписи TS",
	ErrorCodeXADESTFailed:                 "XAdES-T: Ошибка проверки подписи",
	ErrorCodeOCSPRespStatMalformedRequest: "Неправильный запрос",
	ErrorCodeOCSPRespStatInternalError:    "Внутренняя ошибка",
	ErrorCodeOCSPRespStatTryLater:         "Попробуйте позже",
	ErrorCodeOCSPRespStatSigRequired:      "Должны подписать запрос",
	ErrorCodeOCSPRespStatUnauthorized:     "Запрос не авторизован",
	ErrorCodeVerifyIssuerSerialV2:         "не удалось проверить IssuerSerialV2 в XAdES",
	ErrorCodeOCSPCheckCertFromResp:        "Ошибка проверки сертификата OCSP-респондера",
	ErrorCodeCRLExpired:                   "CRL-файл просрочен",
	ErrorCodeLibraryNotInitialized:        "Библиотека не инициализирована",
	ErrorCodeEngineLoadErr:                "Ошибка подключения (загрузки) модуля (engine)",
	ErrorCodeParamError:                   "Некорректные входные данные",
	ErrorCodeCertStatusOK:                 "Статус сертификата – валидный. Используется при проверке сертификата по OCSP. (не является ошибкой, делается запись в лог)",
	ErrorCodeCertStatusRevoked:            "Статус сертификата – отозван. Используется при проверке сертификата по OCSP.",
	ErrorCodeCertStatusUnknown:            "Статус сертификата – неизвестен. Используется при проверке сертификата по OCSP. Например, не удалось установить издателя сертификата.",
}
