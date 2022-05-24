package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getCertificatesList(char *certificates, unsigned long *cert_count) {
//     kc_funcs->KC_GetCertificatesList(certificates, cert_count);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCGetCertificatesList Обеспечивает получение списка сертификатов в виде строки и их количество.
func (cli *KCClient) KCGetCertificatesList() (certs string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: %s", err, r)
			} else {
				err = fmt.Errorf("%w: %s", ErrPanic, r)
			}
		}
	}()

	cli.mu.Lock()
	defer cli.mu.Unlock()

	certsLen := 4096

	cCerts := C.malloc(C.ulong(C.sizeof_char * certsLen))

	defer C.free(cCerts)

	count := 40

	rc := int(C.getCertificatesList(
		(*C.char)(cCerts),
		(*C.ulong)(unsafe.Pointer(&count)),
	))

	certs = C.GoString((*C.char)(cCerts))

	return certs, cli.wrapError(rc)
}
