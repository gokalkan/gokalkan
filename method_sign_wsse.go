package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// unsigned long signWSSE(char *alias, int flags, char *inData, int inDataLength, unsigned char *outSign, int *outSignLength, char *signNodeId) {
//     return kc_funcs->SignWSSE(alias, flags, inData, inDataLength, outSign, outSignLength, signNodeId);
// }
import "C"
import (
	"encoding/xml"
	"fmt"
	"strings"
	"unsafe"

	"github.com/google/uuid"
)

const (
	xmlnsSOAP = "http://schemas.xmlsoap.org/soap/envelope/"
	xmlnsWSU  = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"

	replaceKey = "replace-this"
)

// soapEnvelope представляет soap:Envelope
type soapEnvelope struct {
	XMLName xml.Name `xml:"soap:Envelope"`
	SOAP    string   `xml:"xmlns:soap,attr"`
	WSU     string   `xml:"xmlns:wsu,attr"`
	Body    soapBody `xml:"soap:Body"`
}

// soapBody представляет soap:Body
type soapBody struct {
	ID      string `xml:"wsu:Id,attr"`
	Content string `xml:",chardata"`
}

// SignWSSE подписывает документ XML в формате WSSec, который требутся для SmartBridge
func (cli *Client) SignWSSE(data string) (string, error) {
	soapData := wrapWithSoap(data)

	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))
	flag := 0

	inData := C.CString(soapData)
	defer C.free(unsafe.Pointer(inData))
	inDataLength := len(soapData)

	outSignLength := 50000 + inDataLength
	outSign := C.malloc((C.ulong)(C.sizeof_uchar * outSignLength))
	defer C.free(outSign)

	signNodeID := C.CString("")
	defer C.free(unsafe.Pointer(signNodeID))

	cli.mu.Lock()
	defer cli.mu.Unlock()

	rc := (int)(C.signWSSE(
		alias,
		(C.int)(flag),
		inData,
		(C.int)(inDataLength),
		(*C.uchar)(outSign),
		(*C.int)(unsafe.Pointer(&outSignLength)),
		signNodeID,
	))

	signedXML := C.GoString((*C.char)(outSign))

	return signedXML, cli.returnErr(rc)
}

// wrapWithSoap оборачивает XML документ в SOAP формат, а точнее записывает
// в содержимое под тегом soap:Body
func wrapWithSoap(data string) string {
	id := uuid.New().String()
	idFormatted := fmt.Sprintf("id-%s", id)
	envelope := soapEnvelope{
		SOAP: xmlnsSOAP,
		WSU:  xmlnsWSU,
		Body: soapBody{
			ID:      idFormatted,
			Content: replaceKey,
		},
	}
	b, err := xml.Marshal(envelope)
	if err != nil {
		return ""
	}
	envelopeStr := string(b)
	wrapped := strings.Replace(envelopeStr, replaceKey, data, 1)
	return wrapped
}
