package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// unsigned long signXML(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId, char *parentSignNode, char *parentNameSpace) {
//     return kc_funcs->SignXML(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId, parentSignNode, parentNameSpace);
// }
import "C"
import "unsafe"

// SignXML подписывает данные в формате XML
func (cli *Client) SignXML(data string) (string, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flag := 0

	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(data)

	outSignLength := 50000 + inDataLength
	outSign := C.malloc((C.ulong)(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	signNodeID := C.CString("")
	defer C.free(unsafe.Pointer(signNodeID))

	parentSignNode := C.CString("")
	defer C.free(unsafe.Pointer(parentSignNode))

	parentNameSpace := C.CString("")
	defer C.free(unsafe.Pointer(parentNameSpace))

	rc := (int)(C.signXML(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		signNodeID,
		parentSignNode,
		parentNameSpace,
	))
	signedXML := C.GoString((*C.char)(outSign))

	return signedXML, cli.returnErr(rc)
}
