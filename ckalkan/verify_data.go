package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long verifyData(char *alias, int flags, char *inData, int inDataLength, unsigned char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength) {
//    return kc_funcs->VerifyData(alias, flags, inData, inDataLength, inoutSign, inoutSignLength, outData, outDataLen, outVerifyInfo, outVerifyInfoLen, inCertID, outCert, outCertLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	// длина сертификата возвращаемая от проверки
	outCertLength = 64768

	// длина информации сертификата возвращаемая от проверки
	outVerifyInfoLength = 64768

	// длина данных возвращаемая от проверки
	outDataLength = 28000
)

// VerifiedData структура возвращаемая от метода KCVerifyData
type VerifiedData struct {
	Cert []byte
	Info []byte
	Data []byte
}

// VerifyData обеспечивает проверку подписи.
func (cli *Client) VerifyData(inSign, inData, alias string, flag Flag) (result *VerifiedData, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: panic: %s", err, r)
				return
			}

			err = fmt.Errorf("%w: %s", ErrPanic, r)
		}
	}()

	cli.mu.Lock()
	defer cli.mu.Unlock()

	kcAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(kcAlias))

	kcInData := C.CString(inData)
	defer C.free(unsafe.Pointer(kcInData))
	inDataLength := len(inData)

	kcInSign := unsafe.Pointer(C.CString(inSign))
	defer C.free(kcInSign)
	inputSignLength := len(inSign)

	var kcOutData [outDataLength]byte
	kcOutDataLen := outDataLength

	kcOutVerifyInfo := make([]byte, len(inSign), len(inSign))
	//var kcOutVerifyInfo [outVerifyInfoLength]byte
	kcOutVerifyInfoLen := outVerifyInfoLength

	kcInCertID := 0

	var kcOutCert [outCertLength]byte
	kcOutCertLen := outCertLength

	rc := int(C.verifyData(
		kcAlias,
		C.int(flag),
		kcInData,
		C.int(inDataLength),
		(*C.uchar)(kcInSign),
		C.int(inputSignLength),
		(*C.char)(unsafe.Pointer(&kcOutData)),
		(*C.int)(unsafe.Pointer(&kcOutDataLen)),
		(*C.char)(unsafe.Pointer(&kcOutVerifyInfo)),
		(*C.int)(unsafe.Pointer(&kcOutVerifyInfoLen)),
		C.int(kcInCertID),
		(*C.char)(unsafe.Pointer(&kcOutCert)),
		(*C.int)(unsafe.Pointer(&kcOutCertLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return nil, err
	}

	result = &VerifiedData{
		Cert: byteSlice(kcOutCert[:]),
		Info: byteSlice(kcOutVerifyInfo[:]),
		Data: byteSlice(kcOutData[:]),
	}

	return result, nil
}

func byteSlice(content []byte) []byte {
	for i, v := range content {
		if v == 0 {
			return content[:i]
		}
	}
	return content
}
