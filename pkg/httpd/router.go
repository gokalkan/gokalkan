package httpd

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route ...
func Route(ch ChallengeHandler) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/challenge", ch.SendChallenge).Methods("GET")
	router.HandleFunc("/challenge", ch.HandleChallenge).Methods("POST")
	// router.HandleFunc("/challenges", ch.GetChallenges).Methods("GET")
	http.Handle("/", router)
	return router
}
