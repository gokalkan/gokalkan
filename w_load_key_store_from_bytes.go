package gokalkan

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// LoadKeyStoreFromBytes загружает PKCS12.
func (cli *Client) LoadKeyStoreFromBytes(key []byte, password string) (err error) {
	tmpKey, err := os.CreateTemp("", "tmp.key.*.p12")
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	filename := tmpKey.Name()

	defer func() {
		_ = os.Remove(filename)
	}()

	defer func() {
		_ = tmpKey.Close()
	}()

	written, err := io.Copy(tmpKey, bytes.NewReader(key))
	if err != nil {
		return fmt.Errorf("%w: %s", ErrLoadKey, err)
	}

	_ = tmpKey.Close()

	if exp := int64(len(key)); exp != written {
		return fmt.Errorf("%w: expected %d bytes, but written %d bytes", ErrLoadKey, exp, written)
	}

	return cli.kc.KCLoadKeyStore(password, filename, KCStoreTypePKCS12, "")
}
