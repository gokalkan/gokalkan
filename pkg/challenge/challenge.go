package challenge

import (
	"errors"
	"fmt"

	"github.com/Zulbukharov/kalkan-bind/pkg/storage/memory"
	"github.com/google/uuid"
)

// Service ...
type Service interface {
	GenerateChallenge(login string) (string, error) // registerchallenge and return key
	RegisterChallenge(key string) error             // save to the storage key
	HandleChallenge(xml string) error               // accept xml verify key
	GetChallenges() []memory.Challenge
}

// Repository ...
type Repository interface {
	AddKey(key string) error
	GetKeys() []memory.Challenge
	// VerifyKey(key string) error
	// DeleteKey(key string) error
}

// Bridge ...
type Bridge interface {
	VerifyXML(xml string) (string, int)
}

// type Tools interface {

// }

type service struct {
	tR Repository
	bR Bridge
}

// NewService ...
func NewService(r Repository, b Bridge) Service {
	return &service{r, b}
}

func (s *service) buildChallenge(login string) string {
	// uuid:login
	id := uuid.New()
	return fmt.Sprintf("<challenge>%s:%s</challenge>", id.String(), login)
}

// RegisterChallenge ...
// will add new challenge with "uuid:login"
// and ttl = 1hour
func (s *service) RegisterChallenge(login string) error {
	if err := s.tR.AddKey(login); err != nil {
		return err
	}
	return nil
}

// GenerateChallenge accepts login to sign
// will register challenge and return xml to sign
func (s *service) GenerateChallenge(login string) (string, error) {
	challenge := s.buildChallenge(login)
	err := s.RegisterChallenge(challenge)
	if err != nil {
		return "", err
	}
	return challenge, nil
}

// HandleChallenge accepts signed xml with data, signature, cert
// retuns nil on success and error defined in bridge.consts
func (s *service) HandleChallenge(xml string) error {
	// parse xml => extract original data with challenge tag
	m, rv := s.bR.VerifyXML(xml)
	fmt.Println("HandleChallenge", m, rv)
	if rv != 0 {
		return errors.New(m)
	}
	return nil
}

func (s *service) GetChallenges() []memory.Challenge {
	return s.tR.GetKeys()
}
