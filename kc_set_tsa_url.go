package gokalkan

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

// KCTSASetURL установка адреса сервиса TSA. (Значение по умолчанию http://tsp.pki.gov.kz:80)
func (cli *KCClient) KCTSASetURL(url string) {
	cli.mu.Lock()
	defer cli.mu.Unlock()

	cTSA := C.CString(url)

	defer C.free(unsafe.Pointer(cTSA))

	C.setTSAUrl(
		cTSA,
	)
}
