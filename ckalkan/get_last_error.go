package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getLastError() {
//     return kc_funcs->KC_GetLastError();
// }
import "C"

// GetLastError возвращает код последней ошибки.
func (cli *Client) GetLastError() ErrorCode {
	return ErrorCode(C.getLastError())
}
