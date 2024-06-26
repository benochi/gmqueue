package main

import (
	"fmt"
	"sync"
)

// handle putting in any arbitrary data into storage
// fetch will return data and error
type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

// data slice of bytes.
type MemoryStore struct {
	mu   sync.RWMutex
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, b)
	return len(s.data) - 1, nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}
	return s.data[offset], nil
}
