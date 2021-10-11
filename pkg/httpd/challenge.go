package httpd

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/challenge"
	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/httpd/response"
)

// ChallengeHandler ...
type ChallengeHandler interface {
	HandleChallenge(w http.ResponseWriter, r *http.Request)
	SendChallenge(w http.ResponseWriter, r *http.Request)
	GetChallenges(w http.ResponseWriter, r *http.Request)
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
		return
	}
	response.Execute(w, map[string]string{"status": "ok"}, http.StatusOK, nil)
}

func (c *challengeHandler) SendChallenge(w http.ResponseWriter, r *http.Request) {
	challenge, err := c.cS.GenerateChallenge("user")
	if err != nil {
		response.Execute(w, nil, http.StatusBadRequest, err)
		return
	}
	response.Execute(w, map[string]string{"challenge": challenge}, http.StatusOK, nil)
}

func (c *challengeHandler) GetChallenges(w http.ResponseWriter, r *http.Request) {
	response.Execute(w, c.cS.GetChallenges(), http.StatusOK, nil)
}
