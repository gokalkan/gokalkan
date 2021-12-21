package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// void getLastErrorString(char *errorString, int *bufSize) {
//     kc_funcs->KC_GetLastErrorString(errorString, bufSize);
// }
import "C"
import "unsafe"

const ErrLength = 65534

// GetLastErrorString возвращает текст последней ошибки
func (cli *Client) GetLastErrorString() string {
	errLen := ErrLength
	var errStr [ErrLength]byte

	C.getLastErrorString(
		(*C.char)(unsafe.Pointer(&errStr)),
		(*C.int)(unsafe.Pointer(&errLen)),
	)
	return string(errStr[:])
}
