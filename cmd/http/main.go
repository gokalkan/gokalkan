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

	bridgeService, e := bridge.NewKalkanBridge()
	if e != nil {
		fmt.Println("here?", e)
		return
	}
	defer bridgeService.Close()
	bridgeService.Init()
	bridgeService.KCLoadKeyStore(conf.DigitalSignaturePass, conf.DigitalSignaturePath)

	storageRepo := memory.NewStorage(conf.TTL)
	challengeS := challenge.NewService(storageRepo, bridgeService)

	// generate demo challenge
	z, err := challengeS.GenerateChallenge("IIN012345678901")
	if err != nil {
		fmt.Println(err)
	}
	s, rv := bridgeService.SignXML(z)
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
