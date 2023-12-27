package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long zipConVerify(char *inZipFile, int flags, char *outVerifyInfo, int *outVerifyInfoLen) {
// 		return kc_funcs->ZipConVerify(inZipFile, flags, outVerifyInfo, outVerifyInfoLen);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// ZipConVerify обеспечивает проверку подписи .zip архива.
func (cli *Client) ZipConVerify(inZipFile string, flags Flag) (result string, err error) {
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

	cInZipFile := C.CString(inZipFile)
	defer C.free(unsafe.Pointer(cInZipFile))

	outVerifyInfoLen := 2048
	outVerifyInfo := C.malloc(C.ulong(C.sizeof_uchar * outVerifyInfoLen))
	defer C.free(outVerifyInfo)

	rc := int(C.zipConVerify(
		cInZipFile,
		C.int(int(flags)),
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
