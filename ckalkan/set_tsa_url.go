package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// void setTSAUrl(char *tsaurl) {
//     kc_funcs->KC_TSASetUrl(tsaurl);
// }
import "C"
import (
	"unsafe"
)

// TSASetURL установка адреса сервиса TSA.
func (cli *Client) TSASetURL(url string) {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	cTSA := C.CString(url)
	defer C.free(unsafe.Pointer(cTSA))

	C.setTSAUrl(
		cTSA,
	)
}
