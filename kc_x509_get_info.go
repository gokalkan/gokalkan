package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509CertificateGetInfo(char *inCert, int inCertLength, int propId, char *outData, int *outDataLength) {
//     return kc_funcs->X509CertificateGetInfo(inCert, inCertLength, propId, (unsigned char*)outData, outDataLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

func (cli *KCClient) KCX509CertificateGetInfo(inCert string, prop KCCertProp) (result string, err error) {
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

	outDataLength := 32768

	data := C.malloc(C.ulong(C.sizeof_char * outDataLength))

	defer C.free(data)

	rc := int(C.x509CertificateGetInfo(
		cInCert,
		C.int(len(inCert)),
		C.int(int(prop)),
		(*C.char)(data),
		(*C.int)(unsafe.Pointer(&outDataLength)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return result, err
	}

	result = C.GoString((*C.char)(data))

	return result, nil
}
