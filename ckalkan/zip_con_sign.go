package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long zipConSign(char *alias, const char *filePath, const char *name, const char *outDir, int flags) {
// 		return kc_funcs->ZipConSign(alias, filePath, name, outDir, flags);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// Варианты подписывания:
// 	1. ZipConSign подписывает файлы с последующим размещением их в zip-контейнер.
// 	2. ZipConSign подписывает zip архив (множественная подпись)
// 	3. ZipConSign подписывает все файлы в папке
//
// Параметры:
// 	- alias - label (alias) сертификата.
// 	- filePath
//
//  	1. пути к файлам, которые необходимо записать.
//     	(в конце каждого пути к файлу необходимо вставить вертикальную линию - «|»;
//     	например: inFiles= “./test/1.txt|./test/2.txt|./test/3.txt|”).
//
//  	2. путь к zip архиву
//    	(в конце пути к zip архиву необходимо вставить вертикальную линию - «|»;
//     	например: inFiles= “./test.zip|”).
//
//  	3. путь к папке, в которой лежат файлы на подписание
//     	например: inFiles= “./test/dir”).
//
// 	- name - имя создаваемого архива.
// 	- outDir - расположение создаваемого архива.
// 	- flags - флаги
func (cli *Client) ZipConSign(alias, filePath, name, outDir string, flags Flag) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if err != nil {
				err = fmt.Errorf("%w: panic: %s", err, r)
				return
			}

			err = fmt.Errorf("%w: %s", ErrPanic, r)
		}
	}()

	cli.mu.Lock()
	defer cli.mu.Unlock()

	cAlias := C.CString(alias)
	defer C.free(unsafe.Pointer(cAlias))

	cFilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cFilePath))

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	cOutDir := C.CString(outDir)
	defer C.free(unsafe.Pointer(cOutDir))

	rc := int(C.zipConSign(
		cAlias,
		cFilePath,
		cName,
		cOutDir,
		C.int(int(flags)),
	))

	err = cli.wrapError(rc)
	if err != nil {
		return err
	}

	return nil

}
