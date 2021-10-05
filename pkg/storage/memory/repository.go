package memory

// Storage ...
type Storage struct {
	Keys []string
}

// AddKey ..
func (s *Storage) AddKey(key string) error {
	s.Keys = append(s.Keys, key)
	return nil
}
