package http

import "net/http"

type ChallengeHandler interface {
	HandleChallenge(w http.ResponseWriter, r *http.Request)
	SendChallenge(w http.ResponseWriter, r *http.Request)
}

type challengeHandler struct {
}

func NewChallengeHandler() ChallengeHandler {
	return &challengeHandler{}
}

func (c *challengeHandler) HandleChallenge(w http.ResponseWriter, r *http.Request) {
	//
}

func (c *challengeHandler) SendChallenge(w http.ResponseWriter, r *http.Request) {
	//
}
