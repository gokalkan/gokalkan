package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509ExportCertificateFromStore(char *alias, int flag, char *outCert, int *outCertLength) {
//     return kc_funcs->X509ExportCertificateFromStore(alias, flag, outCert, outCertLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCX509ExportCertificateFromStore экспортирует сертификата из хранилища
func (cli *KCClient) KCX509ExportCertificateFromStore(alias string) (result string, err error) {
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

	flag := 0
	outCertLen := 32768

	cert := C.malloc(C.ulong(C.sizeof_char * outCertLen))

	defer C.free(cert)

	cAlias := C.CString(alias)

	defer C.free(unsafe.Pointer(cAlias))

	rc := int(C.x509ExportCertificateFromStore(
		cAlias,
		C.int(flag),
		(*C.char)(cert),
		(*C.int)(unsafe.Pointer(&outCertLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return result, err
	}

	result = C.GoString((*C.char)(cert))

	return result, nil
}
