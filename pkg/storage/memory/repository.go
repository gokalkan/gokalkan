package memory

import (
	"errors"
	"time"
)

// Storage ...
type Storage interface {
	AddKey(key, serial string) error
	GetChallenge(key string) (*Challenge, error)
	GetKeys() map[string]Challenge
	DeleteChallenge(key string) error
}

// Challenge ...
type Challenge struct {
	Serial    string    `json:"serial"`
	ExpiresAt time.Time `json:"created_at"`
}

type storage struct {
	challenges map[string]Challenge
	// ttl is value in seconds that describes challenge lifetime
	ttl int
}

const (
	challengeDoesNotExist = "Challenge does not exist"
)

// NewStorage ...
func NewStorage(ttl int) Storage {
	return &storage{challenges: make(map[string]Challenge), ttl: ttl}
}

// AddKey ..
func (s *storage) AddKey(key, serial string) error {
	expiresAt := (time.Now().UTC()).Add(time.Second * time.Duration(s.ttl))
	s.challenges[key] = Challenge{serial, expiresAt}
	return nil
}

// GetKeys ...
func (s *storage) GetKeys() map[string]Challenge {
	return s.challenges
}

// GetChallenge ...
func (s *storage) GetChallenge(key string) (*Challenge, error) {
	c, ok := s.challenges[key]
	if !ok {
		return nil, errors.New(challengeDoesNotExist)
	}
	return &c, nil
}

// DeleteChallenge ...
func (s *storage) DeleteChallenge(key string) error {
	_, ok := s.challenges[key]
	if !ok {
		return errors.New(challengeDoesNotExist)
	}
	delete(s.challenges, key)
	return nil
}
