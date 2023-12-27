package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long setProxy(int flags, char *inProxyAddr, char *inProxyPort, char *inUser, char *inPass) {
//     return kc_funcs->KC_SetProxy(flags, inProxyAddr, inProxyPort, inUser, inPass);
// }
import "C"
import (
	"fmt"
	"net/url"
	"unsafe"
)

// SetProxy устанавливает прокси.
func (cli *Client) SetProxy(flags Flag, proxyURL *url.URL) (err error) {
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

	if proxyURL == nil {
		proxyURL = &url.URL{}
	}

	cProxyAddr := C.CString(proxyURL.Hostname())
	defer C.free(unsafe.Pointer(cProxyAddr))

	cProxyPort := C.CString(proxyURL.Port())
	defer C.free(unsafe.Pointer(cProxyPort))

	cProxyUser := C.CString(proxyURL.User.Username())
	defer C.free(unsafe.Pointer(cProxyUser))

	pass, _ := proxyURL.User.Password()
	cProxyPass := C.CString(pass)
	defer C.free(unsafe.Pointer(cProxyPass))

	rc := int(C.setProxy(
		C.int(int(flags)),
		cProxyAddr,
		cProxyPort,
		cProxyUser,
		cProxyPass,
	))

	return cli.wrapError(rc)
}
