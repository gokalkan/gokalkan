package gokalkan

type KCCertProp int // Значение поля/расширения в запросе/сертификате

// Константы, определяющие значение поля/расширения в запросе/сертификате.
const (
	KCCertPropIssuerCountryName   KCCertProp = 2049 // Страна издателя
	KCCertPropIssuerSOPN          KCCertProp = 2050 // Название штата или провинции издателя
	KCCertPropIssuerLocalityName  KCCertProp = 2051 // Населённый пункт издателя
	KCCertPropIssuerOrgName       KCCertProp = 2052 // Наименование организации издателя
	KCCertPropIssuerOrgUnitName   KCCertProp = 2053 // Название организационного подразделения издателя
	KCCertPropIssuerCommonName    KCCertProp = 2054 // Имя Фамилия издателя
	KCCertPropSubjectCountryName  KCCertProp = 2055 // Страна субъекта
	KCCertPropSubjectSOPN         KCCertProp = 2056 // Название штата или провинции субъекта
	KCCertPropSubjectLocalityName KCCertProp = 2057 // Населенный пункт субъекта
	KCCertPropSubjectCommonName   KCCertProp = 2058 // Общее имя субъекта
	KCCertPropSubjectGivenName    KCCertProp = 2059 // Имя субъекта
	KCCertPropSubjectSurname      KCCertProp = 2060 // Фамилия субъекта
	KCCertPropSubjectSerialNumber KCCertProp = 2061 // Серийный номер субъекта
	KCCertPropSubjectEmail        KCCertProp = 2062 // e-mail субъекта
	KCCertPropSubjectOrgName      KCCertProp = 2063 // Наименование организации субъекта
	KCCertPropSubjectOrgUnitName  KCCertProp = 2064 // Название организационного подразделения субъекта
	KCCertPropSubjectBc           KCCertProp = 2065 // Бизнес категория субъекта
	KCCertPropSubjectDc           KCCertProp = 2066 // Доменный компонент субъекта
	KCCertPropNotBefore           KCCertProp = 2067 // Дата действителен с
	KCCertPropNotAfter            KCCertProp = 2068 // Дата действителен по
	KCCertPropKeyUsage            KCCertProp = 2069 // Использование ключа
	KCCertPropExtKeyUsage         KCCertProp = 2070 // Расширенное использование ключа
	KCCertPropAuthKeyID           KCCertProp = 2071 // Идентификатор ключа центра сертификации
	KCCertPropSubjKeyID           KCCertProp = 2072 // Идентификатор ключа субъекта
	KCCertPropCertCN              KCCertProp = 2073 // Серийный номер сертификата
	KCCertPropIssuerDN            KCCertProp = 2074 // Отличительное имя издателя
	KCCertPropSubjectDN           KCCertProp = 2075 // Отличительное имя субъекта
	KCCertPropSignatureAlg        KCCertProp = 2076 // Алгоритм подписи
	KCCertPropPubKey              KCCertProp = 2077 // Получение открытого ключа
	KCCertPropPoliciesID          KCCertProp = 2078 // Получение идентификатора политики сертификата
	KCCertPropOCSP                KCCertProp = 2079 // Получение URL-адреса OCSP
	KCCertPropGetCRL              KCCertProp = 2080 // Получение URL-адреса CRL
	KCCertPropGetDeltaCRL         KCCertProp = 2081 // Получение URL-адреса delta CRL
)

//nolint:gochecknoglobals
var allProps = []KCCertProp{
	KCCertPropIssuerCountryName,
	KCCertPropIssuerSOPN,
	KCCertPropIssuerLocalityName,
	KCCertPropIssuerOrgName,
	KCCertPropIssuerOrgUnitName,
	KCCertPropIssuerCommonName,
	KCCertPropSubjectCountryName,
	KCCertPropSubjectSOPN,
	KCCertPropSubjectLocalityName,
	KCCertPropSubjectCommonName,
	KCCertPropSubjectGivenName,
	KCCertPropSubjectSurname,
	KCCertPropSubjectSerialNumber,
	KCCertPropSubjectEmail,
	KCCertPropSubjectOrgName,
	KCCertPropSubjectOrgUnitName,
	KCCertPropSubjectBc,
	KCCertPropSubjectDc,
	KCCertPropNotBefore,
	KCCertPropNotAfter,
	KCCertPropKeyUsage,
	KCCertPropExtKeyUsage,
	KCCertPropAuthKeyID,
	KCCertPropSubjKeyID,
	KCCertPropCertCN,
	KCCertPropIssuerDN,
	KCCertPropSubjectDN,
	KCCertPropSignatureAlg,
	KCCertPropPubKey,
	KCCertPropPoliciesID,
	KCCertPropOCSP,
	KCCertPropGetCRL,
	KCCertPropGetDeltaCRL,
}
