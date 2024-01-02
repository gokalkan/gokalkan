package ckalkan

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

func (cli *Client) X509LoadCertificateFromBuffer(inCert []byte, certCodetype CertCodeType) (err error) {
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

	rc := int(C.x509LoadCertificateFromBuffer(
		(*C.char)(unsafe.Pointer(&inCert)),
		C.int(len(inCert)),
		C.int(int(certCodetype)),
	))

	return cli.wrapError(rc)
}
