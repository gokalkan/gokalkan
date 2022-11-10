package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getLastErrorString(char *errorString, int *bufSize) {
//     return kc_funcs->KC_GetLastErrorString(errorString, bufSize);
// }
import "C"
import (
	"unsafe"
)

// длина строки с ошибкой возвращаемая от KalkanCrypt
const errLength = 65534

// GetLastErrorString возвращает текст последней ошибки.
func (cli *Client) GetLastErrorString() (code ErrorCode, message string) {
	errLen := errLength

	var errStr [errLength]byte

	rc := int64(C.getLastErrorString(
		(*C.char)(unsafe.Pointer(&errStr)),
		(*C.int)(unsafe.Pointer(&errLen)),
	))

	return ErrorCode(rc), string(byteSlice(errStr[:]))
}
