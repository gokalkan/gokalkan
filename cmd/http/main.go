package main

import (
	"fmt"

	"github.com/Zulbukharov/kalkan-bind/pkg/bridge"
	"github.com/Zulbukharov/kalkan-bind/pkg/challenge"
	"github.com/Zulbukharov/kalkan-bind/pkg/settings"
	"github.com/Zulbukharov/kalkan-bind/pkg/storage/memory"
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
	b.Init()

	m := &memory.Storage{}
	challengeS := challenge.NewService(m, b)

	b.KC_LoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)
	// b.X509ExportCertificateFromStore()
	s, rv := b.SignXML(`<company-id>770704034</company-id>`)
	fmt.Println("SignXML", rv, s)
	// s, rv = b.VerifyXML(s)
	challengeS.HandleChallenge(s)
	// fmt.Println("VerifyXML", rv, s)
	// fmt.Println("VerifyXML", )
	// fmt.Println("VerifyData", b.KC_GetLastErrorString())
	b.Close()
}
