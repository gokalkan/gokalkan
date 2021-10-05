package httpd

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Zulbukharov/kalkan-bind/pkg/challenge"
	"github.com/Zulbukharov/kalkan-bind/pkg/httpd/response"
)

// ChallengeHandler ...
type ChallengeHandler interface {
	HandleChallenge(w http.ResponseWriter, r *http.Request)
	SendChallenge(w http.ResponseWriter, r *http.Request)
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
	Xml string `json: "xml"`
}

func (c *challengeHandler) HandleChallenge(w http.ResponseWriter, r *http.Request) {
	var signedXML SignedXML

	if err := json.NewDecoder(r.Body).Decode(&signedXML); err != nil {
		log.Println("Failed to decode input body: ", err.Error())
		response.Execute(w, nil, http.StatusBadRequest, errors.New(FailedToParse))
		return
	}

	err := c.cS.HandleChallenge(signedXML.Xml)
	if err != nil {
		log.Println("Failed to handle challenge: ", err.Error())
		response.Execute(w, nil, http.StatusBadRequest, errors.New(FailedToParse))
	}
	response.Execute(w, map[string]string{"status": "ok"}, http.StatusOK, nil)
}

func (c *challengeHandler) SendChallenge(w http.ResponseWriter, r *http.Request) {
	c.cS.GenerateChallenge("user")
}
