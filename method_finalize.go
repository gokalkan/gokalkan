package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// void finalize() {
//     return (kc_funcs)->KC_Finalize();
// }
import "C"

// Finalize освобождает ресурсы криптопровайдера KalkanCryptCOM и завершает работу библиотеки
func (cli *Client) Finalize() {
	C.finalize()
}
