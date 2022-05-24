package wtest

import (
	"fmt"
	gokalkan "github.com/Zulbukharov/GoKalkan"
	"os"
	"testing"

	"github.com/Zulbukharov/GoKalkan/internal/testdata/wtest/certs"
)

//nolint:gochecknoglobals
var (
	cli *gokalkan.Client

	key = certs.TestKeyGOST1
)

func TestMain(m *testing.M) {
	var err error

	opts := gokalkan.OptsTest

	cli, err = gokalkan.NewClient(opts...)
	if err != nil {
		fmt.Println("new kalkan client create error", err)
		os.Exit(1)
	}

	err = cli.LoadKeyStoreFromBytes(key.Key, key.Password)
	if err != nil {
		fmt.Println("load key store error", err)
		os.Exit(1)
	}

	c := m.Run()

	er := cli.Close()
	if er != nil {
		fmt.Println("kalkan client close error", er)
	}

	os.Exit(c)
}
