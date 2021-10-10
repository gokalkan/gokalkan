package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/Zulbukharov/kalkan-bind/pkg/bridge"
	"github.com/Zulbukharov/kalkan-bind/pkg/challenge"
	"github.com/Zulbukharov/kalkan-bind/pkg/httpd"
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
	defer b.Close()
	fmt.Println(conf.DigitalSignaturePath)
	b.KCLoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)

	m := memory.NewStorage()
	challengeS := challenge.NewService(m, b)

	// b.X509ExportCertificateFromStore()
	z, err := challengeS.GenerateChallenge("IIN012345678901")
	if err != nil {
		fmt.Println(err)
	}
	s, rv := b.SignXML(z)
	fmt.Println("SignXML", rv, s)

	challengeHandler := httpd.NewChallengeHandler(challengeS)

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", conf.Port),
		Handler: httpd.Route(
			challengeHandler,
		),
	}
	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed ! ")
	}

	// challengeS.HandleChallenge(s)
}
