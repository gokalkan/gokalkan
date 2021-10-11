package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/bridge"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/challenge"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/httpd"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/settings"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/storage/memory"
)

type Message struct {
	XML string `json:"xml"`
}

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

	z, err := challengeS.GenerateChallenge("IIN012345678901")
	if err != nil {
		fmt.Println(err)
	}
	s, rv := b.SignXML(z)
	fmt.Println("SignXML", rv)
	d, _ := json.Marshal(Message{s})
	fmt.Println(string(d))

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
