package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// void finalize() {
//     kc_funcs->KC_Finalize();
// }
import "C"
import "fmt"

// Finalize освобождает ресурсы криптопровайдера KalkanCryptCOM и завершает работу библиотеки
func (cli *Client) Finalize() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nFinalize panic recover: %s\n", r)
		}
	}()

	C.finalize()
}
