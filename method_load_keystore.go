package kalkan

// #cgo LDFLAGS: -ldl
// #include <dlfcn.h>
// #include "cpp/KalkanCrypt.h"
//
// int LoadKeyStore(int storage, char *password, int passLen, char *container, int containerLen, char *alias) {
//     return kc_funcs->KC_LoadKeyStore(storage, password, passLen, container, containerLen, alias);
// }
import "C"
import "unsafe"

// LoadKeyStore загружает ключи/сертификат из хранилища
func (cli *Client) LoadKeyStore(password, containerPath string) error {
	storage := 1 // KCST_PKCS12

	Cpassword := C.CString(password)
	defer C.free(unsafe.Pointer(Cpassword))

	Ccontainer := C.CString(containerPath)
	defer C.free(unsafe.Pointer(Ccontainer))

	alias := C.CString("")
	defer C.free(unsafe.Pointer(alias))

	rc := (int)(C.LoadKeyStore(
		(C.int)(storage),
		Cpassword,
		(C.int)(len(password)),
		Ccontainer,
		(C.int)(len(containerPath)),
		alias,
	))

	return cli.returnErr(rc)
}
