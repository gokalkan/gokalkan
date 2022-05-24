package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509LoadCertificateFromBuffer(char *inCert, int certLength, int flag) {
//     return kc_funcs->X509LoadCertificateFromBuffer((unsigned char*)inCert, certLength, flag);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func (cli *KCClient) KCX509LoadCertificateFromBuffer(inCert string, flag KCCertCodeType) (err error) {
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

	cInCert := C.CString(inCert)

	defer C.free(unsafe.Pointer(cInCert))

	rc := int(C.x509LoadCertificateFromBuffer(
		cInCert,
		C.int(len(inCert)),
		C.int(int(flag)),
	))

	return cli.wrapError(rc)
}
