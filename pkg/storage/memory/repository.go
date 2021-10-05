package memory

type Storage struct {
	Keys []string
}

func (s *Storage) AddKey(key string) {
	s.Keys = append(s.Keys, key)
}
