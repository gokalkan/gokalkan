package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <malloc.h>
//
// void force_defragment() {
//     malloc_trim(0);
// }
import "C"
import (
	"runtime"
	"sync"
)

func getLibraryName() string {
	switch runtime.GOOS {
	case "freebsd":
		return "libkalkancryptwr-64.so.2"
	case "linux":
		return "libkalkancryptwr-64.so"
	default:
		panic("GOOS=" + runtime.GOOS + " is not supported")
	}
}

// требуемая библиотека для KC
var dynamicLibs = getLibraryName()

// Client структура для взаимодействия с библиотекой KC
type Client struct {
	handler *libHandle
	mu      sync.Mutex
}

// NewKCClient возвращает клиента для работы с KC.
func NewClient() (*Client, error) {
	handler, err := getHandle(dynamicLibs)
	if err != nil {
		return nil, err
	}

	cli := &Client{
		handler: handler,
		mu:      sync.Mutex{},
	}

	return cli, nil
}

// forceDefragmentation принудительно дефрагментирует C heap
// Должно вызываться после каждой операции с C.malloc
func (cli *Client) forceDefragmentation() {
	// будет работать только под linux, как и сама либа
	C.force_defragment()

	// Принудительный сбор мусора Go для освобождения heap
	runtime.GC()
}
