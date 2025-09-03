package models

import "fmt"

type InMemoryStore struct {
	items map[string]string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		items: make(map[string]string),
	}
}

func (s *InMemoryStore) Add(item string) error {
	s.items[item] = item
	return nil
}

func (s *InMemoryStore) Get(item string) (string, error) {
	val, ok := s.items[item]
	if !ok {
		return "", fmt.Errorf("item not found")
	}
	return val, nil
}
