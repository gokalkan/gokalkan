package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/doodocs/doodocs/pkg/gokalkan/ckalkan"
	"github.com/doodocs/doodocs/pkg/gokalkan/internal/testdata"
	"github.com/doodocs/doodocs/pkg/gokalkan/internal/testdata/test/certs"
)

//nolint:gochecknoglobals
var (
	cli *ckalkan.Client

	keys = []testdata.Key{
		certs.TestKeyGOST1,
	}
)

func TestMain(m *testing.M) {
	var err error

	cli, err = ckalkan.NewClient()
	if err != nil {
		fmt.Println("new kalkan client create error", err)
		os.Exit(1)
	}

	err = cli.Init()
	if err != nil {
		fmt.Println("new kalkan init error", err)
		os.Exit(1)
	}

	for _, v := range keys {
		err = cli.LoadKeyStore(v.Password, v.Path, ckalkan.StoreTypePKCS12, v.Alias)
		if err != nil {
			fmt.Println("load key store error:", err)
			os.Exit(1)
		}
	}

	c := m.Run()

	er := cli.Close()
	if er != nil {
		fmt.Println("kalkan client close error", er)
	}

	os.Exit(c)
}
