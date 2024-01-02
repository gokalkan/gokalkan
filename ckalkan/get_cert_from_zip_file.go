package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getCertFromZipFile(char* inZipFile, int flags, int inSignID, char *outCert, int *outCertLength) {
//  	return kc_funcs->KC_getCertFromZipFile(inZipFile, flags, inSignID, outCert, outCertLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// GetCertFromZipFile обеспечивает получение сертификата из .zip архива.
func (cli *Client) GetCertFromZipFile(zipFile string, flags Flag, signID int) (cert string, err error) {
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

	cZipFile := C.CString(zipFile)
	defer C.free(unsafe.Pointer(cZipFile))

	outCertLen := 50000
	outCert := C.malloc(C.ulong(C.sizeof_uchar * outCertLen))
	defer C.free(outCert)

	rc := int(C.getCertFromZipFile(
		cZipFile,
		C.int(int(flags)),
		C.int(signID),
		(*C.char)(outCert),
		(*C.int)(unsafe.Pointer(&outCertLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return cert, err
	}

	cert = C.GoString((*C.char)(outCert))

	return cert, nil

}
