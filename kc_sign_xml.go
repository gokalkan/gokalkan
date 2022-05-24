package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long signXML(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace) {
//     return kc_funcs->SignXML(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId, parentSignNode, parentNameSpace);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCSignXML подписывает данные в формате XML.
//  Параметры:
//  xml - входящий xml.
//  alias - label (alias) сертификата.
//  flags - флаги.
//  signNodeID - идентификатор тэга, который необходимо подписать. Передать "", если необходимо подписать все содержимое документа.
//  parentSignNode - идентификатор тэга, в который необходимо поместить значение подписи.
//  parentNameSpace - пространство имен тэга, в который необходимо поместить значение подписи. Если пространство имен есть, но не будет указано - то тег не найдется.
func (cli *KCClient) KCSignXML(xml, alias string, flags KCFlag, signNodeID, parentSignNode, parentNameSpace string) (signedXML string, err error) {
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

	cParentSignNode := C.CString(parentSignNode)

	defer C.free(unsafe.Pointer(cParentSignNode))

	cParentNameSpace := C.CString(parentNameSpace)

	defer C.free(unsafe.Pointer(cParentNameSpace))

	rc := int(C.signXML(
		cAlias,
		C.int(flags),
		cInData,
		C.int(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		cSignNodeID,
		cParentSignNode,
		cParentNameSpace,
	))

	err = cli.wrapError(rc)
	if err != nil {
		return signedXML, err
	}

	signedXML = C.GoString((*C.char)(outSign))

	return signedXML, nil
}
