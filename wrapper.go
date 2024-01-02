package gokalkan

import (
	"errors"
	"fmt"
	"sync"

	"github.com/gokalkan/gokalkan/ckalkan"
	"github.com/gokalkan/gokalkan/types"
)

type Client struct {
	log Logger
	kc  *ckalkan.Client
	o   Options
	mu  sync.Mutex
}

var _ types.Kalkan = (*Client)(nil)

var (
	ErrInit    = errors.New("unable to refer to KC_GetFunctionList")
	ErrHTTPCli = errors.New("http cli error")
)

// NewClient возвращает клиента для работы с KC.
func NewClient(opts ...Option) (*Client, error) {
	o := Options{log: defaultLogger}
	o.setDefaults()

	for _, op := range opts {
		op(&o)
	}

	kc, err := ckalkan.NewClient()
	if err != nil {
		return nil, err
	}

	cli := &Client{
		log: o.log,
		kc:  kc,
		o:   o,
	}

	err = cli.kc.Init()
	if err != nil {
		cli.log.Error("kc init error: ", err)
		return nil, fmt.Errorf("%w: %s", ErrInit, err)
	}

	cli.kc.TSASetURL(cli.o.TSP)

	if cli.o.LoadCerts {
		if err := cli.LoadCerts(); err != nil {
			cli.log.Error("load remote CA certs error: ", err)
			return nil, err
		}
	}

	return cli, nil
}
