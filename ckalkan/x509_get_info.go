package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509CertificateGetInfo(char *inCert, int inCertLength, int propId, char *outData, int *outDataLength) {
//     return kc_funcs->X509CertificateGetInfo(inCert, inCertLength, propId, (unsigned char*)outData, outDataLength);
// }
import "C"
import (
	"bytes"
	"fmt"
	"unsafe"
)

func (cli *Client) X509CertificateGetInfo(inCert string, prop CertProp) (result string, err error) {
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

	outDataLength := C.int(50000)
	outData := make([]byte, outDataLength)

	rc := int(
		C.x509CertificateGetInfo(
			cInCert,
			C.int(len(inCert)),
			C.int(prop),
			(*C.char)(unsafe.Pointer(&outData[0])),
			&outDataLength,
		),
	)

	if err = cli.wrapError(rc); err != nil {
		return "", err
	}

	return string(bytes.Trim(outData[:outDataLength], "\x00")), nil
}
