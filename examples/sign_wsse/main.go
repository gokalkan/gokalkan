package main

import (
	_ "embed"
	"fmt"
	"log"

	kalkan "github.com/Zulbukharov/GoKalkan"
)

var (
	certPath = "test_cert/GOSTKNCA.p12"
	//go:embed test_cert/password.txt
	certPassword string
)

func main() {
	cli, err := kalkan.NewClient()
	if err != nil {
		log.Fatal("NewClient: ", err)
	}
	defer cli.Close()

	if err := cli.LoadKeyStore(certPassword, certPath); err != nil {
		log.Fatal("cli.LoadKeyStore: ", err)
	}

	xml, err := cli.SignWSSE("<root>GoKalkan</root>")
	fmt.Println("xml", xml)
	fmt.Println("err", err)
}
