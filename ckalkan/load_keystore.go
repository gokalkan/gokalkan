package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int loadKeyStore(int storage, char *password, int passLen, char *container, int containerLen, char *alias) {
//     return kc_funcs->KC_LoadKeyStore(storage, password, passLen, container, containerLen, alias);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// LoadKeyStore загружает ключи/сертификат из хранилища
func (cli *Client) LoadKeyStore(password, containerPath string, storeType StoreType, alias string) (err error) {
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

	cPassword := C.CString(password)
	defer C.free(unsafe.Pointer(cPassword))

	cContainer := C.CString(containerPath)
	defer C.free(unsafe.Pointer(cContainer))

	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	rc := int(C.loadKeyStore(
		C.int(int(storeType)),
		cPassword,
		C.int(len(password)),
		cContainer,
		C.int(len(containerPath)),
		cAlias,
	))

	return cli.wrapError(rc)
}
