package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getSigAlgFromXML(const char *xml_in, int xml_in_size, char *retSigAlg, int *retLen) {
//     return kc_funcs->KC_getSigAlgFromXML(xml_in, xml_in_size, retSigAlg, retLen);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// GetSigAlgFromXML обеспечивает получение алгоритма подписи из XML.
func (cli *Client) GetSigAlgFromXML(xml_in string) (sigAlg string, err error) {
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

	cXML := C.CString(xml_in)
	defer C.free(unsafe.Pointer(cXML))

	retLen := 50000
	retSigAlg := C.malloc(C.ulong(C.sizeof_uchar * retLen))
	defer C.free(retSigAlg)

	rc := int(C.getSigAlgFromXML(
		cXML,
		C.int(len(xml_in)),
		(*C.char)(retSigAlg),
		(*C.int)(unsafe.Pointer(&retLen)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return sigAlg, err
	}

	sigAlg = C.GoString((*C.char)(retSigAlg))

	return sigAlg, nil
}
