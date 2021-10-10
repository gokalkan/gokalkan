package challenge

import (
	"errors"
	"fmt"

	"github.com/Zulbukharov/kalkan-bind/pkg/storage/memory"
)

// Service ...
type Service interface {
	GenerateChallenge(login string) (string, error) // registerchallenge and return key
	HandleChallenge(xml string) error               // accept xml verify key
	GetChallenges() map[string]memory.Challenge
}

// Repository ...
type Repository interface {
	AddKey(key, serial string) error
	GetKeys() map[string]memory.Challenge
	VerifyKey(key string) error
	// VerifyKey(key string) error
	// DeleteKey(key string) error
}

// Bridge ...
type Bridge interface {
	VerifyXML(xml string) (string, int)
}

type service struct {
	tR Repository
	bR Bridge
}

// NewService ...
func NewService(r Repository, b Bridge) Service {
	return &service{r, b}
}

// GenerateChallenge accepts login to sign
// will add new challenge
func (s *service) GenerateChallenge(login string) (string, error) {
	challenge := NewChallenge(login, "xml")
	err := s.tR.AddKey(challenge.GetUUID(), challenge.GetSerial())
	if err != nil {
		return "", err
	}
	return challenge.BuildChallenge(), nil
}

// HandleChallenge accepts signed xml with data, signature, cert
// retuns nil on success and error defined in bridge.consts
func (s *service) HandleChallenge(xml string) error {
	// parse xml, validate it, extract challenge uuid
	// verify xml by kalkan, receive serialnumber
	// compare sn, validate challenge expiration, remove it at the end
	challenge, err := ValidateSign([]byte(xml))
	if err != nil {
		return err
	}

	m, rv := s.bR.VerifyXML(xml)
	fmt.Println("HandleChallenge", m, rv)
	if rv != 0 {
		return errors.New(m)
	}

	fmt.Println("Serial from bridge", m)
	err = s.tR.VerifyKey(challenge)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetChallenges() map[string]memory.Challenge {
	return s.tR.GetKeys()
}
