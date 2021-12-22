package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// unsigned long verifyXML(char *alias, int flags, char *inData, int inDataLength, char *outVerifyInfo, int *outVerifyInfoLen) {
// 	   return kc_funcs->VerifyXML(alias, flags, inData, inDataLength, outVerifyInfo, outVerifyInfoLen);
// }
import "C"
import (
	"regexp"
	"strings"
	"unsafe"
)

var serialNumRegExp = regexp.MustCompile(`serialNumber=.*`)

// VerifyXML обеспечивает проверку подписи данных в формате XML
func (cli *Client) VerifyXML(xml string) (string, error) {
	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	flags := 0

	inData := C.CString(xml)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(xml)

	outVerifyInfoLen := 64768
	outVerifyInfo := C.malloc((C.ulong)(C.sizeof_char * outVerifyInfoLen))
	defer C.free(outVerifyInfo)

	cli.mu.Lock()
	defer cli.mu.Unlock()

	rc := (int)(C.verifyXML(
		alias,
		(C.int)(flags),
		inData,
		(C.int)(inDataLength),
		(*C.char)(outVerifyInfo),
		(*C.int)(unsafe.Pointer(&outVerifyInfoLen)),
	))
	if err := cli.returnErr(rc); err != nil {
		return "", err
	}

	outInfo := C.GoString((*C.char)(outVerifyInfo))
	serialNumber := extractSerialNumber(outInfo)
	return serialNumber, nil
}

func extractSerialNumber(info string) string {
	f := serialNumRegExp.FindAllString(info, 1)
	if len(f) == 1 {
		return strings.Replace(f[0], "serialNumber=", "", 1)
	}
	return ""
}
