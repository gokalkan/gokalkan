package main

import (
	"fmt"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/bridge"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/settings"
)

func main() {
	conf, err := settings.ParseYAML("config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	b, e := bridge.NewKalkanBridge()
	if e != nil {
		fmt.Println("here?", e)
		return
	}
	defer b.Close()
	b.Init()
	b.KCLoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)
	s, rv := b.SignXML(`<company-id>770704034</company-id>`)
	fmt.Println("SignXML", s, rv)
	s, rv = b.VerifyXML(s)
	fmt.Println("VerifyXML", s, rv)
}
