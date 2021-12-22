package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// unsigned long verifyData(char *alias, int flags, char *inData, int inDataLength, char *inoutSign, int inoutSignLength, char *outData, int *outDataLen, char *outVerifyInfo, int *outVerifyInfoLen, int inCertID, char *outCert, int *outCertLength) {
//    return kc_funcs->VerifyData(alias, flags, inData, inDataLength, (unsigned char*)inoutSign, inoutSignLength, outData, outDataLen, outVerifyInfo, outVerifyInfoLen, inCertID, outCert, outCertLength);
// }
import "C"
import "unsafe"

const (
	// OutCertLength длина сертификата возвращаемая от проверки
	OutCertLength = 64768

	// OutVerifyInfoLength длина информации сертификата возвращаемая от проверки
	OutVerifyInfoLength = 64768

	// OutVerifyDataLength длина данных возвращаемая от проверки
	OutVerifyDataLength = 28000
)

// VerifiedData структура возвращаемая от метода VerifyData
type VerifiedData struct {
	Cert string
	Info string
	Data string
}

// VerifyData обеспечивает проверку подписи
func (cli *Client) VerifyData(data string) (*VerifiedData, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flag := 0

	inData := C.CString(data)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := 0

	inoutSign := C.CString("")
	defer C.free(unsafe.Pointer(inoutSign))
	inoutSignLength := 0

	var outVerifyData [OutVerifyDataLength]byte
	outVerifyDataLen := OutVerifyDataLength

	var outVerifyInfo [OutVerifyInfoLength]byte
	outVerifyInfoLen := OutVerifyInfoLength

	inCertID := 0

	var outCert [OutCertLength]byte
	outCertLen := OutCertLength

	cli.mu.Lock()
	defer cli.mu.Unlock()

	rc := (int)(C.verifyData(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		inoutSign,
		(C.int)(inoutSignLength),
		(*C.char)(unsafe.Pointer(&outVerifyData)),
		(*C.int)(unsafe.Pointer(&outVerifyDataLen)),
		(*C.char)(unsafe.Pointer(&outVerifyInfo)),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
		(C.int)(inCertID),
		(*C.char)(unsafe.Pointer(&outCert)),
		(*C.int)(unsafe.Pointer(&outCertLen)),
	))
	if err := cli.returnErr(rc); err != nil {
		return nil, err
	}
	return &VerifiedData{
		Cert: string(outCert[:]),
		Info: string(outVerifyInfo[:]),
		Data: string(outVerifyData[:]),
	}, nil
}
