package kalkan

import (
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/dlopen"
)

// dynamicLibs is a list of required libs for Kalkan
var dynamicLibs = []string{"libkalkancryptwr-64.so"}

// Client структура для взаимодействия с библиотекой Kalkan
type Client struct {
	handler *dlopen.LibHandle
}

// NewClient возвращает клиента для работы с Kalkan
func NewClient() (Kalkan, error) {
	handler, err := dlopen.GetHandle(dynamicLibs)
	if err != nil {
		return &Client{}, err
	}

	cli := &Client{
		handler: handler,
	}

	if err := cli.Init(); err != nil {
		return nil, err
	}

	return cli, nil
}