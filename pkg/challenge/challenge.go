package challenge

import "fmt"

/*
	create token by login => save it in memeory

	GET /challenge => returns token
	POST /challenge => verify xml
*/

type Service interface {
	GenerateChallenge() string        // registerchallenge and return key
	RegisterChallenge(key string)     // save to the storage key
	HandleChallenge(xml string) error // accept xml verify key
}

type Repository interface {
	AddKey(key string)
}

type Bridge interface {
	VerifyXML(xml string) (string, int)
}

type service struct {
	tR Repository
	bR Bridge
}

func NewService(r Repository, b Bridge) Service {
	return &service{r, b}
}

func (s *service) RegisterChallenge(key string) {
	s.tR.AddKey(key)
}

func (s *service) GenerateChallenge() string {
	challenge := "<root>ok</root>"
	// generete random xml
	s.RegisterChallenge("<root>ok</root>")
	return challenge
}

func (s *service) HandleChallenge(xml string) error {
	m, rv := s.bR.VerifyXML(xml)
	fmt.Println("HandleChallenge", m, rv)
	return nil
}
