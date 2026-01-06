package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long getCertFromCMS(char *inCMS, int inCMSLen, int inSignId, int flags, char *outCert, int *outCertLength) {
//     return kc_funcs->KC_GetCertFromCMS(inCMS, inCMSLen, inSignId, flags, outCert, outCertLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// GetCertFromCMS обеспечивает получение сертификата из CMS.
func (cli *Client) GetCertFromCMS(cms string, signID int, flag Flag) (cert string, err error) {
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
	defer cli.forceDefragmentation() // Принудительная дефрагментация после операции

	cCMS := C.CString(cms)
	defer C.free(unsafe.Pointer(cCMS))

	outCertLen := 32768
	outCert := C.malloc(C.ulong(C.sizeof_uchar * outCertLen))
	defer C.free(outCert)

	rc := int(C.getCertFromCMS(
		cCMS,
		C.int(len(cms)),
		C.int(signID),
		C.int(int(flag)),
		(*C.char)(outCert),
		(*C.int)(unsafe.Pointer(&outCertLen)),
	))

	// Всегда копируем сертификат ДО проверки ошибки
	if outCert != nil {
		cert = C.GoString((*C.char)(outCert))
	}

	err = cli.wrapError(rc)
	if err != nil {
		// Возвращаем пустую строку при ошибке
		return "", err
	}

	return cert, nil
}
