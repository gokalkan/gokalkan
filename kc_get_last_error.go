package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getLastError() {
//     return kc_funcs->KC_GetLastError();
// }
import "C"

// KCGetLastError возвращает код последней ошибки
func (cli *KCClient) KCGetLastError() KCErrorCode {
	return KCErrorCode(C.getLastError())
}
