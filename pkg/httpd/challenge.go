package httpd

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/challenge"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/httpd/response"
)

// ChallengeHandler ...
type ChallengeHandler interface {
	HandleChallenge(w http.ResponseWriter, r *http.Request)
	SendChallenge(w http.ResponseWriter, r *http.Request)
	// GetChallenges(w http.ResponseWriter, r *http.Request)
}

type challengeHandler struct {
	cS challenge.Service
}

// NewChallengeHandler ...
func NewChallengeHandler(cs challenge.Service) ChallengeHandler {
	return &challengeHandler{cs}
}

// SignedXML ...
type SignedXML struct {
	XML string `json:"xml"`
}

// Response ...
type Response struct {
	Message string `json:"message"`
}

func (c *challengeHandler) HandleChallenge(w http.ResponseWriter, r *http.Request) {
	var signedXML SignedXML

	if err := json.NewDecoder(r.Body).Decode(&signedXML); err != nil {
		log.Println("Failed to decode input body: ", err.Error())
		response.Execute(w, nil, http.StatusBadRequest, errors.New(FailedToParse))
		return
	}

	err := c.cS.HandleChallenge(signedXML.XML)
	if err != nil {
		log.Println("Failed to handle challenge: ", err.Error())
		response.Execute(w, nil, http.StatusBadRequest, errors.New(FailedToParse))
		return
	}
	res := &Response{"ok"}
	response.Execute(w, res, http.StatusOK, nil)
}

func (c *challengeHandler) SendChallenge(w http.ResponseWriter, r *http.Request) {
	serial := r.URL.Query().Get("serial")
	if serial == "" || len(serial) > 20 || len(serial) < 5 {
		response.Execute(w, nil, http.StatusBadRequest, fmt.Errorf("Invalid serial param"))
		return
	}
	challenge, err := c.cS.GenerateChallenge(serial)
	if err != nil {
		response.Execute(w, nil, http.StatusBadRequest, err)
		return
	}
	res := &Response{challenge}
	response.Execute(w, res, http.StatusOK, nil)
}

// func (c *challengeHandler) GetChallenges(w http.ResponseWriter, r *http.Request) {
// 	response.Execute(w, c.cS.GetChallenges(), http.StatusOK, nil)
// }
