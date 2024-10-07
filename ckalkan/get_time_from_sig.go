package ckalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// unsigned long GetTimeFromSig(char *inData, int inDataLength, int flags, int inSigId, time_t *outDateTime) {
//     return kc_funcs->KC_GetTimeFromSig(inData, inDataLength, flags, inSigId, outDateTime);
// }
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func (cli *Client) GetTimeFromSig(cms string, sigID int, flag Flag) (
	timestamp time.Time,
	err error,
) {
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

	cCMS := C.CString(cms)
	defer C.free(unsafe.Pointer(cCMS))

	outDateTime := C.time_t(0)

	rc := int(
		C.GetTimeFromSig(
			cCMS,
			C.int(len(cms)),
			C.int(int(flag)),
			C.int(sigID),
			(*C.time_t)(unsafe.Pointer(&outDateTime)),
		),
	)

	err = cli.wrapError(rc)
	if err != nil {
		return time.Time{}, err
	}

	timestamp = time.Unix(int64(outDateTime), 0)

	return timestamp, nil
}
