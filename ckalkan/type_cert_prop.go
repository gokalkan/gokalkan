package ckalkan

// CertProp определяет значение поля/расширения в запросе/сертификате
type CertProp int

// Константы, определяющие значение поля/расширения в запросе/сертификате.
const (
	CertPropIssuerCountryName   CertProp = 2049 // Страна издателя
	CertPropIssuerSOPN          CertProp = 2050 // Название штата или провинции издателя
	CertPropIssuerLocalityName  CertProp = 2051 // Населённый пункт издателя
	CertPropIssuerOrgName       CertProp = 2052 // Наименование организации издателя
	CertPropIssuerOrgUnitName   CertProp = 2053 // Название организационного подразделения издателя
	CertPropIssuerCommonName    CertProp = 2054 // Имя Фамилия издателя
	CertPropSubjectCountryName  CertProp = 2055 // Страна субъекта
	CertPropSubjectSOPN         CertProp = 2056 // Название штата или провинции субъекта
	CertPropSubjectLocalityName CertProp = 2057 // Населенный пункт субъекта
	CertPropSubjectCommonName   CertProp = 2058 // Общее имя субъекта
	CertPropSubjectGivenName    CertProp = 2059 // Имя субъекта
	CertPropSubjectSurname      CertProp = 2060 // Фамилия субъекта
	CertPropSubjectSerialNumber CertProp = 2061 // Серийный номер субъекта
	CertPropSubjectEmail        CertProp = 2062 // e-mail субъекта
	CertPropSubjectOrgName      CertProp = 2063 // Наименование организации субъекта
	CertPropSubjectOrgUnitName  CertProp = 2064 // Название организационного подразделения субъекта
	CertPropSubjectBc           CertProp = 2065 // Бизнес категория субъекта
	CertPropSubjectDc           CertProp = 2066 // Доменный компонент субъекта
	CertPropNotBefore           CertProp = 2067 // Дата действителен с
	CertPropNotAfter            CertProp = 2068 // Дата действителен по
	CertPropKeyUsage            CertProp = 2069 // Использование ключа
	CertPropExtKeyUsage         CertProp = 2070 // Расширенное использование ключа
	CertPropAuthKeyID           CertProp = 2071 // Идентификатор ключа центра сертификации
	CertPropSubjKeyID           CertProp = 2072 // Идентификатор ключа субъекта
	CertPropCertCN              CertProp = 2073 // Серийный номер сертификата
	CertPropIssuerDN            CertProp = 2074 // Отличительное имя издателя
	CertPropSubjectDN           CertProp = 2075 // Отличительное имя субъекта
	CertPropSignatureAlg        CertProp = 2076 // Алгоритм подписи
	CertPropPubKey              CertProp = 2077 // Получение открытого ключа
	CertPropPoliciesID          CertProp = 2078 // Получение идентификатора политики сертификата
	CertPropOCSP                CertProp = 2079 // Получение URL-адреса OCSP
	CertPropGetCRL              CertProp = 2080 // Получение URL-адреса CRL
	CertPropGetDeltaCRL         CertProp = 2081 // Получение URL-адреса delta CRL
)

//nolint:gochecknoglobals
var AllProps = []CertProp{
	CertPropIssuerCountryName,
	CertPropIssuerSOPN,
	CertPropIssuerLocalityName,
	CertPropIssuerOrgName,
	CertPropIssuerOrgUnitName,
	CertPropIssuerCommonName,
	CertPropSubjectCountryName,
	CertPropSubjectSOPN,
	CertPropSubjectLocalityName,
	CertPropSubjectCommonName,
	CertPropSubjectGivenName,
	CertPropSubjectSurname,
	CertPropSubjectSerialNumber,
	CertPropSubjectEmail,
	CertPropSubjectOrgName,
	CertPropSubjectOrgUnitName,
	CertPropSubjectBc,
	CertPropSubjectDc,
	CertPropNotBefore,
	CertPropNotAfter,
	CertPropKeyUsage,
	CertPropExtKeyUsage,
	CertPropAuthKeyID,
	CertPropSubjKeyID,
	CertPropCertCN,
	CertPropIssuerDN,
	CertPropSubjectDN,
	CertPropSignatureAlg,
	CertPropPubKey,
	CertPropPoliciesID,
	CertPropOCSP,
	CertPropGetCRL,
	// KCCertPropGetDeltaCRL,
}
