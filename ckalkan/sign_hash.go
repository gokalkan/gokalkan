package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long signHash(char *alias, int flags, char *inHash, int inHashLength, unsigned char *outSign, int *outSignLength) {
//     return kc_funcs->SignHash(alias, flags, inHash, inHashLength, outSign, outSignLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// SignHash Подписывает входные хэшированные данные.
//
// Параметры:
//   - algo - алгоритм подписывания ("sha256", "Gost34311_95", "GostR3411_2015_512")
//   - inHash - хэшированные данные
//   - flag - флаги
func (cli *Client) SignHash(algo HashAlgo, inHash string, flag Flag) (signedHash string, err error) {
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

	cAlias := C.CString(string(algo))
	defer C.free(unsafe.Pointer(cAlias))

	cInHash := C.CString(inHash)
	defer C.free(unsafe.Pointer(cInHash))

	inHashLength := len(inHash)
	outSignLength := 50000 + 2*inHashLength
	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	rc := int(C.signHash(
		cAlias,
		C.int(int(flag)),
		cInHash,
		C.int(inHashLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return signedHash, err
	}

	signedHash = C.GoString((*C.char)(outSign))

	return signedHash, nil
}
