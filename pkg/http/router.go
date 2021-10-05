package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Route(ch ChallengeHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/challenge", ch.SendChallenge).Methods("GET")
	router.HandleFunc("/challenge", ch.HandleChallenge).Methods("POST")
	http.Handle("/", router)
	return router
}
