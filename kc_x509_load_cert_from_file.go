package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509LoadCertificateFromFile(char *certPath, int certType) {
//     return kc_funcs->X509LoadCertificateFromFile(certPath, certType);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func (cli *KCClient) KCX509LoadCertificateFromFile(certPath string, certType KCCertType) (err error) {
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

	cCertPath := C.CString(certPath)

	defer C.free(unsafe.Pointer(cCertPath))

	rc := int(C.x509LoadCertificateFromFile(
		cCertPath,
		C.int(int(certType)),
	))

	return cli.wrapError(rc)
}
