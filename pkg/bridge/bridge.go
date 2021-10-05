package bridge

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <string.h>
// #include <stdlib.h>
// #include "KalkanCrypt.h"
// #include <stdio.h>
//
// void printer() {
//    printf("hello\n");
// }
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
	"fmt"
	"unsafe"

	"github.com/Zulbukharov/kalkan-bind/pkg/dlopen"
)

type Kalkan interface {
	Init()
	KC_GetLastErrorString() string
	KC_LoadKeyStore(password, containerPath string)
	X509ExportCertificateFromStore() (string, int)
	VerifyData()
	Close()
	SignXML(data string) (string, int)
	VerifyXML(xml string) (string, int)
}

// unsigned long (*SignData)(char *alias, int flags, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, unsigned char *outSign, int *outSignLength);

type bridge struct {
	handler *dlopen.LibHandle
}

// NewKalkanBridge ...
func NewKalkanBridge() (Kalkan, error) {
	h, e := dlopen.GetHandle([]string{"libkalkancryptwr-64.so"})
	if e != nil {
		return &bridge{}, e
	}
	return &bridge{
		handler: h,
	}, nil
}

func (b *bridge) KC_GetLastErrorString() string {
	errLen := 65534
	var errStr [65534]byte
	C.BindKC_GetLastErrorString((*C.char)(unsafe.Pointer(&errStr)), (*C.int)(unsafe.Pointer(&errLen)))
	return string(errStr[:])
}

func (b *bridge) Init() {
	f, err := b.handler.GetSymbolPointer("KC_GetFunctionList")
	if err != nil {
		fmt.Printf(`couldn't get symbol %v\n`, err)
		return
	}
	C.BindKC_GetFunctionList(f)
	fmt.Printf("%v\n", int(C.BindInit()))
}

func (b *bridge) KC_LoadKeyStore(password, containerPath string) {
	// unsigned long KC_LoadKeyStore(int storage, char *password, int passLen, char *container, int containerLen, char *alias) {}
	storage := 1 // KCST_PKCS12
	Cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(Cpassword))
	Ccontainer := C.CString(containerPath)
	defer C.free(unsafe.Pointer(Ccontainer))

	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rv := C.BindKC_LoadKeyStore((C.int)(storage), Cpassword, (C.int)(len(password)), Ccontainer, (C.int)(len(containerPath)), alias)
	fmt.Println("KC_LoadKeyStore", int(rv))
}

func (b *bridge) X509ExportCertificateFromStore() (string, int) {
	// int (*X509ExportCertificateFromStore)(char *alias, int flag, char *outCert, int *outCertLength)
	flag := 1
	outCertLength := 32768
	cert := C.malloc((C.ulong)(C.sizeof_char * outCertLength))
	defer C.free(cert)
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rv := (int)(C.BindX509ExportCertificateFromStore(
		alias,
		(C.int)(flag),
		(*C.char)(cert),
		(*C.int)(unsafe.Pointer(&outCertLength)),
	))
	if rv != 0 {
		return b.KC_GetLastErrorString(), rv
	}
	return C.GoString((*C.char)(cert)), rv
}

// VerifyData returs verification result of signature by data and public key
func (b *bridge) VerifyData() {
	// accepts inData
	// unsigned long BindVerifyData(char *alias, int flags, char *inData, int inDataLength, char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength)
	/*
		alias = NULL
		flags = 0
		inData // data to verify (hello world)
		inDataLength
		inoutSign // data signature crypted by private key
		inoutSignLength
		outData
		outDataLen
		outVerifyInfo
		outVerifyInfoLen
		inCertID
		outCert
		outCertLength
	*/
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))
	flag := 0
	inData := C.CString("as")
	defer C.free(unsafe.Pointer(inData))
	inDataLength := 0
	inoutSign := C.CString("")
	defer C.free(unsafe.Pointer(inoutSign))
	inoutSignLength := 0
	outDataLen := 28000
	outData := C.malloc((C.ulong)(C.sizeof_char * outDataLen))
	defer C.free(outData)
	var outVerifyInfo [64768]byte
	outVerifyInfoLen := 64768
	inCertID := 0
	var outCert [64768]byte
	outCertLength := 64768
	C.BindVerifyData(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		inoutSign,
		(C.int)(inoutSignLength),
		(*C.char)(outData),
		(*C.int)(unsafe.Pointer(&outDataLen)),
		(*C.char)(unsafe.Pointer(&outVerifyInfo)),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
		(C.int)(inCertID),
		(*C.char)(unsafe.Pointer(&outCert)),
		(*C.int)(unsafe.Pointer(&outCertLength)),
	)
}

func (b *bridge) Close() {
	b.handler.Close()
}

// SignXML returns signed xml and result status
func (b *bridge) SignXML(data string) (string, int) {
	// unsigned long (*SignXML)(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace);
	/*
		alias ""
		flags 0
		inData "ok"
		inDataLength 0
		outSign []{}
		outSignLength 50000 + inDataLength
		signNodeId ""
		parentSignNode ""
		parentNameSpace ""
	*/
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))
	flag := 0
	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(data)
	outSignLength := 50000 + inDataLength
	outSign := C.malloc((C.ulong)(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)
	signNodeId := C.CString("")
	defer C.free(unsafe.Pointer(signNodeId))
	parentSignNode := C.CString("")
	defer C.free(unsafe.Pointer(parentSignNode))
	parentNameSpace := C.CString("")
	defer C.free(unsafe.Pointer(parentNameSpace))
	rv := (int)(C.BindSignXML(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		signNodeId,
		parentSignNode,
		parentNameSpace,
	))
	if rv != 0 {
		return b.KC_GetLastErrorString(), rv
	}
	return C.GoString((*C.char)(outSign)), rv
}

// VerifyXML returns C function return value
func (b *bridge) VerifyXML(xml string) (string, int) {
	// unsigned long (*VerifyXML)(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen);
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))
	flags := 0
	inData := C.CString(xml)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(xml)
	outVerifyInfoLen := 64768
	outVerifyInfo := C.malloc((C.ulong)(C.sizeof_char * outVerifyInfoLen))
	defer C.free(outVerifyInfo)
	rv := (int)(C.BindVerifyXML(
		alias,
		(C.int)(flags),
		inData,
		(C.int)(inDataLength),
		(*C.char)(outVerifyInfo),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
	))
	if rv != 0 {
		return b.KC_GetLastErrorString(), rv
	}
	return C.GoString((*C.char)(outVerifyInfo)), rv
}
