package ckalkan

// #cgo LDFLAGS: -ldl
// #include <stdlib.h>
// #include <dlfcn.h>
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

var ErrHandler = errors.New("lib handler error")

// LibHandle represents an open handle to a library (.so)
type libHandle struct {
	Handle  unsafe.Pointer
	LibName string
}

// GetHandle tries to get a handle to a library (.so), attempting to access it
// by the names specified in libs and returning the first that is successfully
// opened. Callers are responsible for closing the handler. If no library can
// be successfully opened, an error is returned.
func getHandle(name string) (*libHandle, error) {
	libName := C.CString(name)

	defer C.free(unsafe.Pointer(libName))

	handle := C.dlopen(libName, C.RTLD_LAZY)
	if handle != nil {
		h := &libHandle{
			Handle:  handle,
			LibName: name,
		}

		return h, nil
	}

	return nil, fmt.Errorf("%w: %s", ErrHandler, C.GoString(C.dlerror()))
}

// GetSymbolPointer takes a symbol name and returns a pointer to the symbol.
func (l *libHandle) getSymbolPointer(symbol string) (unsafe.Pointer, error) {
	sym := C.CString(symbol)

	defer C.free(unsafe.Pointer(sym))

	C.dlerror()

	p := C.dlsym(l.Handle, sym)

	e := C.dlerror()
	if e != nil {
		return nil, fmt.Errorf("%w: error resolving symbol %q: %s", ErrHandler, symbol, C.GoString(e))
	}

	return p, nil
}

// Close closes a LibHandle.
func (l *libHandle) close() error {
	C.dlerror()
	C.dlclose(l.Handle)

	e := C.dlerror()
	if e != nil {
		return fmt.Errorf("%w: error closing %s: %s", ErrHandler, l.LibName, C.GoString(e))
	}

	return nil
}
