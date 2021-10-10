package memory

import (
	"errors"
	"fmt"
	"time"
)

// Storage ...
type Storage interface {
	AddKey(key, serial string) error
	VerifyKey(key string) error
	GetKeys() map[string]Challenge
}

// Challenge ...
type Challenge struct {
	Serial    string    `json:"serial"`
	CreatedAt time.Time `json:"created_at"`
}

type storage struct {
	challenges map[string]Challenge
}

// NewStorage ...
func NewStorage() Storage {
	return &storage{challenges: make(map[string]Challenge)}
}

// AddKey ..
func (s *storage) AddKey(key, serial string) error {
	s.challenges[key] = Challenge{serial, time.Now().UTC()}
	return nil
}

// GetKeys ...
func (s *storage) GetKeys() map[string]Challenge {
	return s.challenges
}

// VerifyKey ...
func (s *storage) VerifyKey(challenge string) error {
	if val, ok := s.challenges[challenge]; ok {
		fmt.Println("Storage val", val)
		// delete
		return nil
	}
	return errors.New("Challenge does not exist")
}
