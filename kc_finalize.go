package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// void finalize() {
//     kc_funcs->KC_Finalize();
// }
import "C"
import "fmt"

// KCFinalize освобождает ресурсы криптопровайдера KalkanCryptCOM и завершает работу библиотеки
func (cli *KCClient) KCFinalize() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nFinalize panic recover: %s\n", r)
		}
	}()

	C.finalize()
}
