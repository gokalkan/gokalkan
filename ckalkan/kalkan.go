package ckalkan

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
