package gokalkan

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/doodocs/doodocs/pkg/gokalkan/ckalkan"
)

// LoadKeyStore загружает PKCS12.
func (cli *Client) LoadKeyStore(path, password string) error {
	return cli.kc.LoadKeyStore(password, path, ckalkan.StoreTypePKCS12, "")
}

// LoadKeyStoreFromBytes загружает PKCS12.
func (cli *Client) LoadKeyStoreFromBytes(key []byte, password string) (err error) {
	tmpKey, err := os.CreateTemp("", "tmp.key.*.p12")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	filename := tmpKey.Name()

	defer os.Remove(filename)
	defer tmpKey.Close()

	written, err := io.Copy(tmpKey, bytes.NewReader(key))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	if exp := int64(len(key)); exp != written {
		return fmt.Errorf("%w: expected %d bytes, but written %d bytes", ErrLoadKey, exp, written)
	}

	return cli.kc.LoadKeyStore(password, filename, ckalkan.StoreTypePKCS12, "")
}
