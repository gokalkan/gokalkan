package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// 	unsigned long GetTimeFromSig(char *inData, int inDataLength, int flags, int inSigId, time_t *outDateTime) {
//     return kc_funcs->KC_GetTimeFromSig(inData, inDataLength, flags, inSigId, outDateTime);
// }
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

// GetTimeFromSig обеспечивает получение времени подписи из CMS.
//
// *inData - входные данные (подпись, в текущей версии только формата CAdES)*
//
// *flags - должны быть установлены флаги входящего формата(PEM, DER, ...etc)*
func (cli *Client) GetTimeFromSig(inData string, flags Flag, inSigId int) (outDateTime time.Time, err error) {
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

	cData := C.CString(inData)
	defer C.free(unsafe.Pointer(cData))

	var outDateTimeC C.time_t

	rc := int(C.GetTimeFromSig(
		cData,
		C.int(len(inData)),
		C.int(int(flags)),
		C.int(inSigId),
		&outDateTimeC,
	))

	err = cli.wrapError(rc)
	if err != nil {
		return outDateTime, err
	}

	outDateTime = time.Unix(int64(outDateTimeC), 0)

	return outDateTime, nil
}
