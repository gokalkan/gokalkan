package test

import (
	"fmt"
	gokalkan "github.com/Zulbukharov/GoKalkan"
	"os"
	"testing"

	"github.com/Zulbukharov/GoKalkan/internal/testdata"
	"github.com/Zulbukharov/GoKalkan/internal/testdata/test/certs"
)

//nolint:gochecknoglobals
var (
	cli *gokalkan.KCClient

	keys = []testdata.Key{
		certs.TestKeyGOST1,
	}
)

func TestMain(m *testing.M) {
	var err error

	cli, err = gokalkan.NewKCClient()
	if err != nil {
		fmt.Println("new kalkan client create error", err)
		os.Exit(1)
	}

	err = cli.KCInit()
	if err != nil {
		fmt.Println("new kalkan init error", err)
		os.Exit(1)
	}

	for _, v := range keys {
		err = cli.KCLoadKeyStore(v.Password, v.Path, gokalkan.KCStoreTypePKCS12, v.Alias)
		if err != nil {
			fmt.Println("load key store error:", err)
			os.Exit(1)
		}
	}

	c := m.Run()

	er := cli.KCClose()
	if er != nil {
		fmt.Println("kalkan client close error", er)
	}

	os.Exit(c)
}
