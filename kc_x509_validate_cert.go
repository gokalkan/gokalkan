package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "KalkanCrypt.h"
//
// int x509ValidateCertificate(char *inCert, int inCertLength, int validType, char *validPath, long long checkTime, char *outInfo, int *outInfoLength) {
//     return kc_funcs->X509ValidateCertificate(inCert, inCertLength, validType, validPath, checkTime, outInfo, outInfoLength);
// }
import "C"
import (
	"fmt"
	"unsafe"
)

// KCX509ValidateCertificate - осуществляет проверку сертификата: проверка срока действия, построение цепочки сертификатов, проверка отозванности по OCSP или CRL.
//  Если validateType:
//  - KCValidateTypeCRL - в параметр path необходимо указывать путь к файлу crl.
//  Например:
//  X509ValidateCertificate(gostCert, KCValidateTypeCRL, "/tmp/nca_gost.crl")
//
//  - KCValidateTypeOCSP - в параметр path необходимо указывать url OCSP. По умолчанию передается url http://ocsp.pki.gov.kz.
//  Например:
//  X509ValidateCertificate(gostCert, KCValidateTypeOCSP)
//  X509ValidateCertificate(gostCert, KCValidateTypeOCSP, "http://ocsp.pki.gov.kz")
//
//  - KCValidateTypeNothing - не производятся проверки по CRL или OCSP. Параметр path игнорируется.
//  Например:
//  X509ValidateCertificate(gostCert, KCValidateTypeNothing, "")
func (cli *KCClient) KCX509ValidateCertificate(inCert string, validateType KCValidateType, validatePath string) (result string, err error) {
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

	CInCert := C.CString(inCert)

	defer C.free(unsafe.Pointer(CInCert))

	CValidPath := C.CString(validatePath)

	defer C.free(unsafe.Pointer(CValidPath))

	outDataLength := 32768

	data := C.malloc(C.ulong(C.sizeof_char * outDataLength))

	defer C.free(data)

	rc := int(C.x509ValidateCertificate(
		CInCert,
		C.int(len(inCert)),
		C.int(int(validateType)),
		CValidPath,
		0,
		(*C.char)(data),
		(*C.int)(unsafe.Pointer(&outDataLength)),
	))

	return C.GoString((*C.char)(data)), cli.wrapError(rc)
}
