package main

import (
	"fmt"

	"github.com/Zulbukharov/kalkan-bind/pkg/bridge"
	"github.com/Zulbukharov/kalkan-bind/pkg/settings"
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
	// b.X509ExportCertificateFromStore()
	s, rv := b.SignXML(`<company-id>770704034</company-id>`)
	fmt.Println("SignXML", rv)
	s, rv = b.VerifyXML(s)
	fmt.Println("VerifyXML", rv)
	// fmt.Println("VerifyXML", )
	// fmt.Println("VerifyData", b.KC_GetLastErrorString())
}
