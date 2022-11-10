package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdio.h>
// #include "KalkanCrypt.h"
//
// void getFunctionList(void *f) {
//     void (*KC_GetFunctionList)(stKCFunctionsType **);
//     KC_GetFunctionList = (void (*)(stKCFunctionsType **))f;
//     KC_GetFunctionList(&kc_funcs);
// }
//
// int init() {
//     int rv = (kc_funcs)->KC_Init();
//     return rv;
// }
import "C"
import (
	"fmt"
)

// Init инициализирует библиотеку
func (cli *Client) Init() (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: panic: %s", err, r)
				return
			}

			err = fmt.Errorf("%w: %s", ErrPanic, r)
		}
	}()

	f, err := cli.handler.getSymbolPointer("KC_GetFunctionList")
	if err != nil {
		return err
	}

	C.getFunctionList(f)
	rc := int(C.init())

	return cli.wrapError(rc)
}
