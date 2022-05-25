package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long signWSSE(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId) {
//     return kc_funcs->SignWSSE(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCSignWSSE подписывает документ XML в формате WSSec, который требутся для SmartBridge.
//
// Параметры:
//  xml - строка XML.
//  alias - label (alias) сертификата.
//  flags - флаги.
//  signNodeID - идентификатор тэга, который необходимо подписать. Передать "", если необходимо подписать все содержимое документа.
func (cli *KCClient) KCSignWSSE(xml, alias string, flags KCFlag, signNodeID string) (signedXML string, err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: panic: %s", err, r)
				return
			}

			err = fmt.Errorf("%w: %s", ErrPanic, r)
		}
	}()

	cAlias := C.CString(alias)

	defer C.free(unsafe.Pointer(cAlias))

	cInData := C.CString(xml)

	defer C.free(unsafe.Pointer(cInData))

	inDataLength := len(xml)

	outSignLength := 50000 + inDataLength

	outSign := C.malloc(C.ulong(C.sizeof_uchar * outSignLength))

	defer C.free(outSign)

	cSignNodeID := C.CString(signNodeID)

	defer C.free(unsafe.Pointer(cSignNodeID))

	cli.mu.Lock()
	defer cli.mu.Unlock()

	rc := int(C.signWSSE(
		cAlias,
		C.int(flags),
		cInData,
		C.int(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		cSignNodeID,
	))

	err = cli.wrapError(rc)
	if err != nil {
		return signedXML, err
	}

	signedXML = C.GoString((*C.char)(outSign))

	return signedXML, nil
}
