package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// void GetFunctionList(void *f) {
//     void (*KC_GetFunctionList)(stKCFunctionsType **);
//     KC_GetFunctionList = (void (*)(stKCFunctionsType **))f;
//     KC_GetFunctionList(&kc_funcs);
// }
//
// int Init() {
//     return (kc_funcs)->KC_Init();
// }
import "C"
import "errors"

// Init инициализирует библиотеку
func (cli *Client) Init() error {
	f, err := cli.handler.GetSymbolPointer("KC_GetFunctionList")
	if err != nil {
		return errors.New("unable to refer to KC_GetFunctionList - " + err.Error())
	}

	C.GetFunctionList(f)
	rc := int(C.Init())

	return cli.returnErr(rc)
}
