package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getCertFromXML(char *inXML, int inXMLLen, int inSignId, char *outCert, int *outCertLength) {
//     return kc_funcs->KC_getCertFromXML(inXML, inXMLLen, inSignId, outCert, outCertLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// GetCertFromXML обеспечивает получение сертификата из XML.
func (cli *Client) GetCertFromXML(xml string, signID int) (cert string, err error) {
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

	cXML := C.CString(xml)
	defer C.free(unsafe.Pointer(cXML))

	outCertLen := C.int(50000)
	outCert := make([]byte, outCertLen)

	rc := int(
		C.getCertFromXML(
			cXML,
			C.int(len(xml)),
			C.int(signID),
			(*C.char)(unsafe.Pointer(&outCert[0])),
			&outCertLen,
		),
	)

	if err = cli.wrapError(rc); err != nil {
		return "", err
	}

	return string(outCert[:outCertLen]), nil
}
