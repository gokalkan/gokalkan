package gokalkan

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

// errLength длина строки с ошибкой возвращаемая от KalkanCrypt
const errLength = 65534

// KCGetLastErrorString возвращает текст последней ошибки
func (cli *KCClient) KCGetLastErrorString() (code KCErrorCode, message string) {
	errLen := errLength

	var errStr [errLength]byte

	rc := int64(C.getLastErrorString(
		(*C.char)(unsafe.Pointer(&errStr)),
		(*C.int)(unsafe.Pointer(&errLen)),
	))

	return KCErrorCode(rc), string(errStr[:])
}
