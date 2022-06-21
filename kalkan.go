package gokalkan

import (
	"net/url"
	"sync"
)

// KC - интерфейс с методами KalkanCrypt
type KC interface {
	// KCInit инициализирует библиотеку
	KCInit() error

	// KCGetTokens Обеспечивает получение указателя на строку подключенных устройств типа storage и их количество.
	KCGetTokens(store KCStoreType) (tokens string, err error)
	// KCGetCertificatesList Обеспечивает получение списка сертификатов в виде строки и их количество.
	KCGetCertificatesList() (certs string, err error)
	// KCGetCertFromCMS Обеспечивает получение сертификата из CMS.
	KCGetCertFromCMS(cms string, signId int, flag KCFlag) (cert string, err error)
	// KCLoadKeyStore загружает ключи/сертификат из хранилища
	KCLoadKeyStore(password, containerPath string, storeType KCStoreType, alias string) (err error)

	// KCSignXML подписывает данные в формате XML.
	KCSignXML(xml, alias string, flags KCFlag, signNodeID, parentSignNode, parentNameSpace string) (signedXML string, err error)
	// KCVerifyXML обеспечивает проверку подписи данных в формате XML.
	KCVerifyXML(xml, alias string, flags KCFlag) (result string, err error)

	// KCSignWSSE подписывает документ XML в формате WSSec, который требутся для SmartBridge
	KCSignWSSE(xml, alias string, flags KCFlag, signNodeID string) (signedXML string, err error)

	// KCSignData используется для подписи текста в формате base64
	KCSignData(inSign, inData string, alias string, flag KCFlag) (string, error)
	// KCVerifyData обеспечивает проверку подписи
	KCVerifyData(inSign, inData string, alias string, flag KCFlag) (*VerifiedData, error)
	KCHashData(algo KCHashAlgo, dataB64 string, flag KCFlag) (result string, err error)

	KCGetLastError() KCErrorCode
	KCGetLastErrorString() (KCErrorCode, string)

	// KCX509ExportCertificateFromStore экспортирует сертификата из хранилища
	KCX509ExportCertificateFromStore(alias string) (string, error)
	// KCX509ValidateCertificate - осуществляет проверку сертификата: проверка срока действия, построение цепочки сертификатов, проверка отозванности по OCSP или CRL
	KCX509ValidateCertificate(inCert string, validateType KCValidateType, path string) (result string, err error)
	KCX509CertificateGetInfo(inCert string, prop KCCertProp) (result string, err error)
	KCX509LoadCertificateFromBuffer(inCert string, flag KCCertCodeType) error
	KCX509LoadCertificateFromFile(certPath string, certType KCCertType) error

	// KCSetProxy устанавливает прокси.
	KCSetProxy(flag KCFlag, proxyURL *url.URL) error
	// KCTSASetURL установка адреса сервиса TSA. (Значение по умолчанию http://tsp.pki.gov.kz:80)
	KCTSASetURL(url string)
	// KCClose закрывает связь с динамической библиотекой
	KCClose() error
}

// требуемая библиотека для KC
const dynamicLibs = "libkalkancryptwr-64.so"

var _ KC = (*KCClient)(nil)

// KCClient структура для взаимодействия с библиотекой KC
type KCClient struct {
	handler *LibHandle
	mu      sync.Mutex
}

// NewKCClient возвращает клиента для работы с KC.
func NewKCClient() (*KCClient, error) {
	handler, err := GetHandle(dynamicLibs)
	if err != nil {
		return nil, err
	}

	cli := &KCClient{
		handler: handler,
		mu:      sync.Mutex{},
	}

	return cli, nil
}

// wrapError возвращает последнюю глобальную ошибку, если returnCode не равен 0
func (cli *KCClient) wrapError(returnCode int) error {
	if returnCode != 0 {
		ec, es := cli.KCGetLastErrorString()

		return KalkanError{
			errorCode:   ec,
			errorString: es,
		}
	}

	return nil
}
