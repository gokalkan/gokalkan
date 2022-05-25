package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long verifyData(char *alias, int flags, char *inData, int inDataLength, char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength) {
//    return kc_funcs->VerifyData(alias, flags, inData, inDataLength, (unsigned char*)inoutSign, inoutSignLength, outData, outDataLen, outVerifyInfo, outVerifyInfoLen, inCertID, outCert, outCertLength);
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
	outVerifyDataLength = 28000
)

// VerifiedData структура возвращаемая от метода KCVerifyData
type VerifiedData struct {
	Cert string
	Info string
	Data string
}

// KCVerifyData обеспечивает проверку подписи
func (cli *KCClient) KCVerifyData(data, alias string, flag KCFlag) (result *VerifiedData, err error) {
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

	cAlias := C.CString(alias)

	defer C.free(unsafe.Pointer(cAlias))

	inData := C.CString(data)

	defer C.free(unsafe.Pointer(inData))

	inDataLength := 0

	inputSign := C.CString("")

	defer C.free(unsafe.Pointer(inputSign))

	inputSignLength := 0

	var outVerifyData [outVerifyDataLength]byte

	outVerifyDataLen := outVerifyDataLength

	var outVerifyInfo [outVerifyInfoLength]byte

	outVerifyInfoLen := outVerifyInfoLength

	inCertID := 0

	var outCert [outCertLength]byte

	outCertLen := outCertLength

	rc := int(C.verifyData(
		cAlias,
		C.int(flag),
		inData,
		C.int(inDataLength),
		inputSign,
		C.int(inputSignLength),
		(*C.char)(unsafe.Pointer(&outVerifyData)),
		(*C.int)(unsafe.Pointer(&outVerifyDataLen)),
		(*C.char)(unsafe.Pointer(&outVerifyInfo)),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
		C.int(inCertID),
		(*C.char)(unsafe.Pointer(&outCert)),
		(*C.int)(unsafe.Pointer(&outCertLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return nil, err
	}

	result = &VerifiedData{
		Cert: string(outCert[:]),
		Info: string(outVerifyInfo[:]),
		Data: string(outVerifyData[:]),
	}

	return result, nil
}
