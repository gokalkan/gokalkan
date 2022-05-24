package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// void xmlFinalize() {
//     kc_funcs->KC_XMLFinalize();
// }
import "C"
import "fmt"

// KCXMLFinalize освобождает память и завершает работу библиотеки с модулями,
// отвечающие за парсинг, подпись и проверку данных в формате XML.
// Не надо вызывать каждый раз при подписи. Можно только один раз после цикла подписания xml файлов
func (cli *KCClient) KCXMLFinalize() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\nXMLFinalize panic recover: %s\n", r)
		}
	}()

	C.xmlFinalize()
}
