package challenge

import (
	"errors"
	"fmt"
	"time"

	"github.com/Zulbukharov/kalkancrypt-wrapper/pkg/storage/memory"
)

// Service ...
type Service interface {
	GenerateChallenge(login string) (string, error)
	HandleChallenge(xml string) error // accept xml verify key
	GetChallenges() map[string]memory.Challenge
}

// Repository ...
type Repository interface {
	AddKey(key, serial string) error
	GetKeys() map[string]memory.Challenge
	GetChallenge(key string) (*memory.Challenge, error)
	DeleteChallenge(key string) error
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

// HandleChallenge accepts signed xml with data, signature, cert \
// retuns nil on success and error defined in bridge.consts
// parse xml, validate it, extract challenge uuid
// verify xml by kalkan, receive serialnumber
// compare sn, validate challenge expiration, removes it at the end
func (s *service) HandleChallenge(xml string) error {

	challenge, err := ValidateSign([]byte(xml))
	if err != nil {
		return err
	}

	m, rv := s.bR.VerifyXML(xml)
	if rv != 0 {
		return errors.New(m)
	}

	storageChallenge, err := s.tR.GetChallenge(challenge)
	if err != nil {
		return err
	}

	err = s.tR.DeleteChallenge(challenge)
	if err != nil {
		return err
	}

	if storageChallenge.Serial != m {
		return fmt.Errorf("Signature serial is invalid: %s", m)
	}
	if storageChallenge.ExpiresAt.Before(time.Now().UTC()) {
		return fmt.Errorf("Challenge expired")
	}
	return nil
}

func (s *service) GetChallenges() map[string]memory.Challenge {
	return s.tR.GetKeys()
}
