package ckalkan

import (
	"sync"
)

// требуемая библиотека для KC
const dynamicLibs = "libkalkancryptwr-64.so"

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
