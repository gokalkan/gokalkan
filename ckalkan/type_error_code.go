package ckalkan

import "strconv"

type ErrorCode int64 // Код ошибки времени выполнения

func (t ErrorCode) String() string {
	v, ok := errorLabels[t]
	if ok {
		return v
	}

	return "Неизвестный код ошибки"
}

const baseHex = 16

func (t ErrorCode) Hex() string {
	return strconv.FormatInt(int64(t), baseHex)
}

const (
	ErrorCodeOK                           ErrorCode = 0         // Нет ошибки
	ErrorCodeErrorReadPKCS12              ErrorCode = 149946370 // Невозможно прочитать файл формата pkcs#12
	ErrorCodeErrorOpenPKCS12              ErrorCode = 149946371 // Невозможно открыть файл формата pkcs12
	ErrorCodeInvalidPropID                ErrorCode = 149946372 // Недопустимый идентификатор расширения сертификата
	ErrorCodeBufferTooSmall               ErrorCode = 149946373 // Размер буфера слишком мал
	ErrorCodeCertParseError               ErrorCode = 149946374 // Невозможно разобрать (распарсить) сертификат
	ErrorCodeInvalidFlag                  ErrorCode = 149946375 // Недопустимый флаг
	ErrorCodeOpenFileErr                  ErrorCode = 149946376 // Невозможно открыть файл
	ErrorCodeInvalidPassword              ErrorCode = 149946377 // Неправильный пароль
	ErrorCodeMemoryError                  ErrorCode = 149946381 // Невозможно выделить память
	ErrorCodeCheckChainError              ErrorCode = 149946382 // Не найден сертификат УЦ или сертификат пользователя при проверки цепочки
	ErrorCodeValidTypeError               ErrorCode = 149946384 // Недопустимый тип валидации сертификата
	ErrorCodeBadCRLFormat                 ErrorCode = 149946385 // Некорректный формат CRL
	ErrorCodeLoadCRLError                 ErrorCode = 149946386 // Невозможно загрузить CRL
	ErrorCodeLoadCRLsError                ErrorCode = 149946387 // Невозможно загрузить CRL-ы
	ErrorCodeUnknownAlg                   ErrorCode = 149946389 // Неизвестный алгоритм подписи
	ErrorCodeKeyNotFound                  ErrorCode = 149946390 // Не найден приватный ключ пользователя
	ErrorCodeSignInitError                ErrorCode = 149946391 // Невозможно инициализировать менеджера подписи
	ErrorCodeSignError                    ErrorCode = 149946392 // Не удалось сгенерировать цифровую подпись
	ErrorCodeEncodeError                  ErrorCode = 149946393 // Ошибка шифрования
	ErrorCodeInvalidFlags                 ErrorCode = 149946394 // Недопустимые флаги
	ErrorCodeCertNotFound                 ErrorCode = 149946395 // Не найден сертификат пользователя
	ErrorCodeVerifySignError              ErrorCode = 149946396 // Ошибка верификации подписи xml
	ErrorCodeBase64DecodeError            ErrorCode = 149946397 // Ошибка дешифровки из Base 64
	ErrorCodeUnknownCMSFormat             ErrorCode = 149946398 // Неизвестный формат CMS
	ErrorCodeCACertNotFound               ErrorCode = 149946400 // Не найден сертификат УЦ
	ErrorCodeXMLSecInitError              ErrorCode = 149946401 // Ошибка инициализации xmlsec
	ErrorCodeLoadTrustedCertsErr          ErrorCode = 149946402 // Ошибка загрузки доверенных сертификатов
	ErrorCodeSignInvalid                  ErrorCode = 149946403 // Недопустимая подпись xml
	ErrorCodeNoSignFound                  ErrorCode = 149946404 // Не найдена подпись во входных данных
	ErrorCodeDecodeError                  ErrorCode = 149946405 // Ошибка дешифрования
	ErrorCodeXMLParseError                ErrorCode = 149946406 // Невозможно разобрать (распарсить) xml
	ErrorCodeXMLAddIDError                ErrorCode = 149946407 // Не удалось добавить атрибут ID
	ErrorCodeXMLInternalError             ErrorCode = 149946408 // Ошибка при работе с xml
	ErrorCodeXMLSetSignError              ErrorCode = 149946409 // Не удалось подписать xml
	ErrorCodeOpenSSLError                 ErrorCode = 149946410 // Ошибка openssl
	ErrorCodeNoTokenFound                 ErrorCode = 149946412 // Не найден токен
	ErrorCodeOCSPAddCertErr               ErrorCode = 149946413 // Не удалось добавить сертификат в ocsp
	ErrorCodeOCSPParseURLErr              ErrorCode = 149946414 // Не удалось разобрать url
	ErrorCodeOCSPAddHostErr               ErrorCode = 149946415 // Не удалось добавить хост
	ErrorCodeOCSPReqErr                   ErrorCode = 149946416 // Не удалось добавить текущее время в запрос
	ErrorCodeOCSPConnectionErr            ErrorCode = 149946417 // Ошибка подключения к OCSP респондеру
	ErrorCodeVerifyNoData                 ErrorCode = 149946418 // Нет входных данных для верификации
	ErrorCodeIDAttrNotFound               ErrorCode = 149946419 // Не найден атрибут ID
	ErrorCodeIDRange                      ErrorCode = 149946420 // Некорректный идентификатор
	ErrorCodeReaderNotFound               ErrorCode = 149946423 // Не найден ридер
	ErrorCodeGetCertPropErr               ErrorCode = 149946424 // Не удалось получить значение атрибута
	ErrorCodeSignFormat                   ErrorCode = 149946425 // Неизвестный формат подписи
	ErrorCodeInDataFormat                 ErrorCode = 149946426 // Неизвестный формат входных данных
	ErrorCodeOutDataFormat                ErrorCode = 149946427 // Неизвестный формат выходных данных
	ErrorCodeVerifyInitError              ErrorCode = 149946428 // Невозможно инициализировать менеджера верификации подписи
	ErrorCodeVerifyError                  ErrorCode = 149946429 // Не удалось верифицировать цифровую подпись
	ErrorCodeHashError                    ErrorCode = 149946430 // Не удалось хэшировать данные
	ErrorCodeSignHashError                ErrorCode = 149946431 // Не удалось подписать хэшированные данные
	ErrorCodeCACertsNotFound              ErrorCode = 149946432 // Не найден сертификат УЦ в хранилище сертификатов
	ErrorCodeCertTimeInvalid              ErrorCode = 149946434 // Срок действия сертификата истек либо еще не наступил
	ErrorCodeConvertError                 ErrorCode = 149946435 // Ошибка записи сертификата в структуру X509
	ErrorCodeTSACreateQuery               ErrorCode = 149946436 // Ошибка генерации запроса timestamp
	ErrorCodeCreateObj                    ErrorCode = 149946437 // Ошибка записи OID в ASN1 структуру
	ErrorCodeCreateNoNce                  ErrorCode = 149946438 // Ошибка генерации уникального числа
	ErrorCodeHTTPError                    ErrorCode = 149946439 // Ошибка протокола http
	ErrorCodeCADESBESFailed               ErrorCode = 149946440 // Ошибка проверки расширения CADESBES в CMS
	ErrorCodeCADESTFailed                 ErrorCode = 149946441 // Ошибка проверки подписи токена TSA
	ErrorCodeNoTSAToken                   ErrorCode = 149946442 // В подписи не присутствует метка TSA
	ErrorCodeInvalidDigestLen             ErrorCode = 149946443 // Неправильная длина хэша
	ErrorCodeGenRandError                 ErrorCode = 149946444 // Ошибка генерации случайного числа
	ErrorCodeSoapNSError                  ErrorCode = 149946445 // Не найдены заголовки SOAP-сообщений
	ErrorCodeGetPubKey                    ErrorCode = 149946446 // Ошибка экспорта публичного ключа
	ErrorCodeGetCertInfo                  ErrorCode = 149946447 // Ошибка получения информации о сертификате
	ErrorCodeFileReadError                ErrorCode = 149946448 // Ошибка чтения файла
	ErrorCodeCheckError                   ErrorCode = 149946449 // Хэш не совпадает
	ErrorCodeZipExtractErr                ErrorCode = 149946450 // Невозможно открыть архив
	ErrorCodeNoManifestFile               ErrorCode = 149946451 // Не найден MANIFEST
	ErrorCodeVerifyTSHash                 ErrorCode = 149946452 // не удалось проверить Хэш подписи TS
	ErrorCodeXADESTFailed                 ErrorCode = 149946453 // XAdES-T: Ошибка проверки подписи
	ErrorCodeOCSPRespStatMalformedRequest ErrorCode = 149946454 // Неправильный запрос
	ErrorCodeOCSPRespStatInternalError    ErrorCode = 149946455 // Внутренняя ошибка
	ErrorCodeOCSPRespStatTryLater         ErrorCode = 149946456 // Попробуйте позже
	ErrorCodeOCSPRespStatSigRequired      ErrorCode = 149946457 // Должны подписать запрос
	ErrorCodeOCSPRespStatUnauthorized     ErrorCode = 149946458 // Запрос не авторизован
	ErrorCodeVerifyIssuerSerialV2         ErrorCode = 149946459 // не удалось проверить IssuerSerialV2 в XAdES
	ErrorCodeOCSPCheckCertFromResp        ErrorCode = 149946460 // Ошибка проверки сертификата OCSP-респондера
	ErrorCodeCRLExpired                   ErrorCode = 149946461 // CRL-файл просрочен
	ErrorCodeLibraryNotInitialized        ErrorCode = 149946625 // Библиотека не инициализирована
	ErrorCodeEngineLoadErr                ErrorCode = 149946880 // Ошибка подключения (загрузки) модуля (engine)
	ErrorCodeParamError                   ErrorCode = 149947136 // Некорректные входные данные
	ErrorCodeCertStatusOK                 ErrorCode = 149947392 // Статус сертификата – валидный. Используется при проверке сертификата по OCSP. (не является ошибкой, делается запись в лог)
	ErrorCodeCertStatusRevoked            ErrorCode = 149947393 // Статус сертификата – отозван. Используется при проверке сертификата по OCSP.
	ErrorCodeCertStatusUnknown            ErrorCode = 149947394 // Статус сертификата – неизвестен. Используется при проверке сертификата по OCSP. Например, не удалось установить издателя сертификата.
)

var errorLabels = map[ErrorCode]string{
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
