package gokalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include <strings.h>
// #include "KalkanCrypt.h"
//
// int x509ValidateCertificate(char *inCert, int inCertLength, int validType, char *validPath, long long checkTime, char *outInfo, int *outInfoLength) {
//     bzero(outInfo, *outInfoLength);
//     int rc = kc_funcs->X509ValidateCertificate(inCert, inCertLength, validType, validPath, checkTime, outInfo, outInfoLength);
//     for(size_t i = 0; i < *outInfoLength-1; i++)
//         if(outInfo[i] == '\r' && outInfo[i+1] != '\n') outInfo[i] = '\n';
//     return rc;
// }
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	outInfoLength = 8192
)

// KCX509ValidateCertificate - осуществляет проверку сертификата: проверка срока действия,
// построение цепочки сертификатов, проверка отозванности по OCSP или CRL.
// Если validateType:
//
// - KCValidateTypeCRL - в параметр path необходимо указывать путь к файлу crl.
// 	 Например:
//		X509ValidateCertificate(gostCert, KCValidateTypeCRL, "/tmp/nca_gost.crl")
//
// - KCValidateTypeOCSP - в параметр path необходимо указывать url OCSP. По умолчанию передается url http://ocsp.pki.gov.kz.
//   Например:
//		X509ValidateCertificate(gostCert, KCValidateTypeOCSP)
//		X509ValidateCertificate(gostCert, KCValidateTypeOCSP, "http://ocsp.pki.gov.kz")
//
// - KCValidateTypeNothing - не производятся проверки по CRL или OCSP. Параметр path игнорируется.
// 	 Например:
//		X509ValidateCertificate(gostCert, KCValidateTypeNothing, "")
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

	var kcOutInfo [outInfoLength]byte
	kcOutInfoLen := outInfoLength

	rc := int(C.x509ValidateCertificate(
		CInCert,
		C.int(len(inCert)),
		C.int(int(validateType)),
		CValidPath,
		0,
		(*C.char)(unsafe.Pointer(&kcOutInfo)),
		(*C.int)(unsafe.Pointer(&kcOutInfoLen)),
	))

	return string(byteSlice(kcOutInfo[:])), cli.wrapError(rc)
}
