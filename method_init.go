package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <stdio.h>
// #include "cpp/KalkanCrypt.h"
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
	"errors"
)

// Init инициализирует библиотеку
func (cli *Client) Init() error {
	f, err := cli.handler.GetSymbolPointer("KC_GetFunctionList")
	if err != nil {
		return errors.New("unable to refer to KC_GetFunctionList - " + err.Error())
	}
	C.getFunctionList(f)
	rc := int(C.init())

	return cli.returnErr(rc)
}
