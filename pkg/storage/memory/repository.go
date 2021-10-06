package memory

import (
	"errors"
	"time"
)

// Storage ...
type Storage interface {
	AddKey(key string) error
	VerifyKey(key string) error
	GetKeys() []Challenge
}

// Challenge ...
type Challenge struct {
	Challenge string    `json:"challenge"`
	CreatedAt time.Time `json:"created_at"`
}

type storage struct {
	challenges []Challenge
}

// NewStorage ...
func NewStorage() Storage {
	return &storage{challenges: nil}
}

// AddKey ..
func (s *storage) AddKey(challenge string) error {
	newChallenge := Challenge{Challenge: challenge, CreatedAt: time.Now().UTC()}
	s.challenges = append(s.challenges, newChallenge)
	return nil
}

// GetKeys ...
func (s *storage) GetKeys() []Challenge {
	return s.challenges
}

// VerifyKey ...
func (s *storage) VerifyKey(challenge string) error {
	del := -1
	for i, v := range s.challenges {
		if v.Challenge == challenge {
			del = i
			break
		}
	}
	if del == -1 {
		return errors.New("Challenge does not exist")
	}
	s.challenges = append(s.challenges[:del], s.challenges[del+1:]...)
	return nil

}
