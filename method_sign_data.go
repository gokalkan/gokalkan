package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <strings.h>
// #include "cpp/KalkanCrypt.h"
//
// unsigned long signData(char *alias, int flag, char *inData, int inDataLength, unsigned char *inSign, int inSignLen, unsigned char *outSign, int *outSignLength) {
//     bzero(outSign, *outSignLength);
//     return kc_funcs->SignData(alias, flag, inData, inDataLength, inSign, inSignLen, outSign, outSignLength);
// }
import "C"
import (
	"unsafe"
)

const (
	// KC_SIGN_CMS | KC_IN_BASE64 | KC_OUT_BASE6
	SignBase64 = 2066
)

// SignData используется для подписи текста в формате base64
func (cli *Client) SignData(data string) (string, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(data)

	outSignLength := 50000 + 2*inDataLength
	outSign := C.malloc((C.ulong)(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	inSignLength := 50000 + 2*inDataLength
	inSign := C.malloc((C.ulong)(C.sizeof_uchar * inSignLength))
	defer C.free(inSign)

	cli.mu.Lock()
	defer cli.mu.Unlock()

	rc := (int)(C.signData(
		alias,
		(C.int)(SignBase64),
		inData,
		(C.int)(inDataLength),
		(*C.uchar)(inSign),
		(C.int)(inSignLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
	))
	return C.GoString((*C.char)(outSign)), cli.returnErr(rc)
}
