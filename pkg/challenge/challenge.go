package challenge

import (
	"errors"
	"fmt"
)

// Service ...
type Service interface {
	GenerateChallenge(login string) (string, error) // registerchallenge and return key
	RegisterChallenge(key string) (string, error)   // save to the storage key
	HandleChallenge(xml string) error               // accept xml verify key
}

// Repository ...
type Repository interface {
	AddKey(key string) error
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

// RegisterChallenge ...
// will add new challenge with "uuid:login"
// and ttl = 1hour
func (s *service) RegisterChallenge(login string) (string, error) {
	// generate uuid
	if err := s.tR.AddKey(login); err != nil {
		return "", err
	}
	return login, nil
}

// GenerateChallenge accepts login to sign
// will register challenge and return xml to sign
func (s *service) GenerateChallenge(login string) (string, error) {
	key, err := s.RegisterChallenge(login)
	if err != nil {
		return "", err
	}
	challenge := fmt.Sprintf("<challenge>%s</challenge>", key)
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
