package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getTokens(unsigned long storage, char *tokens, unsigned long *tk_count) {
//     return kc_funcs->KC_GetTokens(storage, tokens, tk_count);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCGetTokens обеспечивает получение указателя на строку подключенных устройств типа storage и их количество.
func (cli *KCClient) KCGetTokens(store KCStoreType) (tokens string, err error) {
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

	tokenLen := 8192

	cTokens := C.malloc(C.ulong(C.sizeof_char * tokenLen))

	defer C.free(cTokens)

	count := uint64(0)

	rc := int(C.getTokens(
		C.ulong(uint(store)),
		(*C.char)(cTokens),
		(*C.ulong)(unsafe.Pointer(&count)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return tokens, err
	}

	tokens = C.GoString((*C.char)(cTokens))

	return tokens, nil
}
