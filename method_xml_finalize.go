package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// void XMLFinalize() {
//     return (kc_funcs)->KC_XMLFinalize();
// }
import "C"

// XMLFinalize освобождает память и завершает работу библиотеки с модулями,
// отвечающие за парсинг, подпись и проверку данных в формате XML.
// Не надо вызывать каждый раз при подписи. Можно только один раз после цикла подписания xml файлов
func (cli *Client) XMLFinalize() {
	C.XMLFinalize()
}
