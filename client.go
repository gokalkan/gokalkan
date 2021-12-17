package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <string.h>
// #include <stdlib.h>
// #include "cpp/KalkanCrypt.h"
// #include <stdio.h>
//
// stKCFunctionsType *kc_funcs;
//
// void BindKC_GetFunctionList(void *f) {
//     void (*KC_GetFunctionList)(stKCFunctionsType **);
//     KC_GetFunctionList = (void (*)(stKCFunctionsType **))f;
//     KC_GetFunctionList(&kc_funcs);
// }
//
// int BindInit() {
//     return (kc_funcs)->KC_Init();
// }
//
// int BindKC_LoadKeyStore(int storage, char *password, int passLen, char *container, int containerLen, char *alias) {
//     return kc_funcs->KC_LoadKeyStore(storage, password, passLen, container, containerLen, alias);
// }
//
// int BindX509ExportCertificateFromStore(char *alias, int flag, char *outCert, int *outCertLength) {
//     return kc_funcs->X509ExportCertificateFromStore(alias, flag, outCert, outCertLength);
// }
//
// void BindKC_GetLastErrorString(char *errorString, int *bufSize) {
//     kc_funcs->KC_GetLastErrorString(errorString, bufSize);
// }
//
// unsigned long BindVerifyData(char *alias, int flags, char *inData, int inDataLength, char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength) {
//    return kc_funcs->VerifyData(alias, flags, inData, inDataLength, (unsigned char*)inoutSign, inoutSignLength, outData, outDataLen, outVerifyInfo, outVerifyInfoLen, inCertID, outCert, outCertLength);
// }
//
// unsigned long BindSignXML(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace) {
//     return kc_funcs->SignXML(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId, parentSignNode, parentNameSpace);
// }
//
// unsigned long BindVerifyXML(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen) {
// 	   return kc_funcs->VerifyXML(alias, flags, inData, inDataLength, outVerifyInfo, outVerifyInfoLen);
// }
import "C"

import (
	"errors"
	"regexp"
	"strings"
	"unsafe"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/dlopen"
)

// dynamicLibs is a list of required libs for Kalkan
var dynamicLibs = []string{"libkalkancryptwr-64.so"}

// Client структура для взаимодействия с библиотекой Kalkan
type Client struct {
	handler *dlopen.LibHandle
}

// NewClient возвращает клиента для работы с Kalkan
func NewClient() (Kalkan, error) {
	handler, err := dlopen.GetHandle(dynamicLibs)
	if err != nil {
		return &Client{}, err
	}

	cli := &Client{
		handler: handler,
	}

	if err := cli.Init(); err != nil {
		return nil, err
	}

	return cli, nil
}

// GetLastErrorString возвращает текст последней ошибки
func (cli *Client) GetLastErrorString() string {
	errLen := 65534
	errStr := make([]byte, errLen)

	C.BindKC_GetLastErrorString(
		(*C.char)(unsafe.Pointer(&errStr)),
		(*C.int)(unsafe.Pointer(&errLen)),
	)

	return string(errStr[:])
}

// Init инициализирует библиотеку
func (cli *Client) Init() error {
	f, err := cli.handler.GetSymbolPointer("KC_GetFunctionList")
	if err != nil {
		return errors.New("unable to refer to KC_GetFunctionList - " + err.Error())
	}

	C.BindKC_GetFunctionList(f)
	rc := int(C.BindInit())

	return cli.returnErr(rc)
}

// LoadKeyStore загружает ключи/сертификат из хранилища
func (cli *Client) LoadKeyStore(password, containerPath string) error {
	storage := 1 // KCST_PKCS12

	Cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(Cpassword))

	Ccontainer := C.CString(containerPath)
	defer C.free(unsafe.Pointer(Ccontainer))

	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rc := (int)(C.BindKC_LoadKeyStore(
		(C.int)(storage),
		Cpassword,
		(C.int)(len(password)),
		Ccontainer,
		(C.int)(len(containerPath)),
		alias,
	))

	return cli.returnErr(rc)
}

// X509ExportCertificateFromStore экспортирует сертификата из хранилища
func (cli *Client) X509ExportCertificateFromStore() (string, error) {
	flag := 1
	outCertLength := 32768

	cert := C.malloc((C.ulong)(C.sizeof_char * outCertLength))
	defer C.free(cert)
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rc := (int)(C.BindX509ExportCertificateFromStore(
		alias,
		(C.int)(flag),
		(*C.char)(cert),
		(*C.int)(unsafe.Pointer(&outCertLength)),
	))
	resultCert := C.GoString((*C.char)(cert))

	return resultCert, cli.returnErr(rc)
}

// VerifiedData структура возвращаемая от метода VerifyData
type VerifiedData struct {
	Cert string
	Info string
	Data string
}

// VerifyData обеспечивает проверку подписи
// TODO: принимать аргументы и возврашать результат и ошибку
func (cli *Client) VerifyData(data string) (*VerifiedData, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flag := 0

	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := 0

	inoutSign := C.CString("")
	defer C.free(unsafe.Pointer(inoutSign))
	inoutSignLength := 0

	outVerifyDataLen := 28000
	outVerifyData := make([]byte, outVerifyDataLen)

	outVerifyInfoLen := 64768
	outVerifyInfo := make([]byte, outVerifyInfoLen)

	inCertID := 0

	outCertLength := 64768
	outCert := make([]byte, outCertLength)

	rc := (int)(C.BindVerifyData(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		inoutSign,
		(C.int)(inoutSignLength),
		(*C.char)(unsafe.Pointer(&outVerifyData)),
		(*C.int)(unsafe.Pointer(&outVerifyDataLen)),
		(*C.char)(unsafe.Pointer(&outVerifyInfo)),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
		(C.int)(inCertID),
		(*C.char)(unsafe.Pointer(&outCert)),
		(*C.int)(unsafe.Pointer(&outCertLength)),
	))
	if err := cli.returnErr(rc); err != nil {
		return nil, err
	}
	return &VerifiedData{
		Cert: string(outCert[:]),
		Info: string(outVerifyInfo[:]),
		Data: string(outVerifyData[:]),
	}, nil
}

// Close закрывает связь с динамической библиотекой
func (cli *Client) Close() error {
	return cli.handler.Close()
}

// SignXML подписывает данные в формате XML
func (cli *Client) SignXML(data string) (string, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flag := 0

	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(data)

	outSignLength := 50000 + inDataLength
	outSign := C.malloc((C.ulong)(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	signNodeID := C.CString("")
	defer C.free(unsafe.Pointer(signNodeID))

	parentSignNode := C.CString("")
	defer C.free(unsafe.Pointer(parentSignNode))

	parentNameSpace := C.CString("")
	defer C.free(unsafe.Pointer(parentNameSpace))

	rc := (int)(C.BindSignXML(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		signNodeID,
		parentSignNode,
		parentNameSpace,
	))
	signedXML := C.GoString((*C.char)(outSign))

	return signedXML, cli.returnErr(rc)
}

// VerifyXML обеспечивает проверку подписи данных в формате XML
func (cli *Client) VerifyXML(xml string) (string, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flags := 0

	inData := C.CString(xml)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(xml)

	outVerifyInfoLen := 64768
	outVerifyInfo := C.malloc((C.ulong)(C.sizeof_char * outVerifyInfoLen))
	defer C.free(outVerifyInfo)

	rc := (int)(C.BindVerifyXML(
		alias,
		(C.int)(flags),
		inData,
		(C.int)(inDataLength),
		(*C.char)(outVerifyInfo),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
	))
	outInfo := C.GoString((*C.char)(outVerifyInfo))
	serialNumber := extractSerialNumber(outInfo)

	return serialNumber, cli.returnErr(rc)
}

// returnErr возвращает последнюю глобальную ошибку, если returnCode не равен 0
func (cli *Client) returnErr(returnCode int) error {
	if returnCode != 0 {
		return errors.New(cli.GetLastErrorString())
	}
	return nil
}

func extractSerialNumber(info string) string {
	re := regexp.MustCompile(`serialNumber=.*`)
	f := re.FindAllString(info, 1)
	if len(f) == 1 {
		return strings.Replace(f[0], "serialNumber=", "", 1)
	}
	return ""
}
