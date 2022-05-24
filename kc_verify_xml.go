package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long verifyXML(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen) {
// 	   return kc_funcs->VerifyXML(alias, flags, inData, inDataLength, outVerifyInfo, outVerifyInfoLen);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCVerifyXML обеспечивает проверку подписи данных в формате XML.
func (cli *KCClient) KCVerifyXML(xml, alias string, flags KCFlag) (result string, err error) {
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

	inData := C.CString(xml)

	defer C.free(unsafe.Pointer(inData))

	inDataLength := len(xml)

	outVerifyInfoLen := 64768

	outVerifyInfo := C.malloc(C.ulong(C.sizeof_char * outVerifyInfoLen))

	defer C.free(outVerifyInfo)

	rc := int(C.verifyXML(
		cAlias,
		C.int(flags),
		inData,
		C.int(inDataLength),
		(*C.char)(outVerifyInfo),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return result, err
	}

	result = C.GoString((*C.char)(outVerifyInfo))

	return result, nil
}
