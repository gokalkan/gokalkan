package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <strings.h>
// #include "KalkanCrypt.h"
//
// unsigned long signData(char *alias, int flag, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, unsigned char *outSign, int *outSignLength) {
//     bzero(outSign, *outSignLength);
//     return kc_funcs->SignData(alias, flag, inData, inDataLength, inSign, inSignLen, outSign, outSignLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	TestBase64        = KCFlagInBase64 | KCFlagSignCMS
	SignBase64        = KCFlagSignCMS | KCFlagInBase64 | KCFlagOutBase64
	SignBase64WithTSP = KCFlagSignCMS | KCFlagInBase64 | KCFlagOutBase64 | KCFlagWithTimestamp
)

// KCSignData используется для подписи текста в формате base64
func (cli *KCClient) KCSignData(data, alias string, flag KCFlag) (result string, err error) {
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

	inDataLength := len(data)

	outSignLength := 50000 + 2*inDataLength

	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))

	defer C.free(outSign)

	inSignLength := 50000 + 2*inDataLength

	inSign := C.malloc(C.ulong(C.sizeof_uchar * inSignLength))

	defer C.free(inSign)

	rc := int(C.signData(
		cAlias,
		C.int(int(flag)),
		inData,
		C.int(inDataLength),
		(*C.uchar)(inSign),
		C.int(inSignLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
	))

	return C.GoString((*C.char)(outSign)), cli.wrapError(rc)
}
