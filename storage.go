package main

import "fmt"

//handle putting in any arbitrary data into storage
//fetch will return data and error
type Storer interface {
	Push([]byte) (int, error)
	Fetch(int) ([]byte, error)
}

//data slice of bytes.
type MemoryStore struct {
	data [][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make([][]byte, 0),
	}
}

func (s *MemoryStore) Push(b []byte) (int, error) {
	s.data = append(s.data, b)
	return len(s.data), nil
}

func (s *MemoryStore) Fetch(offset int) ([]byte, error) {
	if len(s.data) < offset {
		return nil, fmt.Errorf("offset (%d) too high", offset)
	}
	return s.data[offset], nil
}
