package main

import (
	"fmt"
	"slices"
)

type Store struct {
	data map[string]string
}

// Keys returns every key currently in the store, sorted alphabetically.
func (s *Store) Keys() []string {
	out := make([]string, 0, len(s.data))
	for k := range s.data {
		out = append(out, k)
	}
	slices.Sort(out)
	return out
}

func (s *Store) Get(key string) (string, bool) {
	value, exists := s.data[key]
	if exists {
		return value, exists
	}
	return "", false
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {
	store := NewStore()
	store.Set("a", "42")
	store.Set("b", "72")
	store.Delete("a")

	_, exists := store.Get("a")

	fmt.Println(exists)

	fmt.Println("GoKV - Go key value store project")
}
