package main

import "fmt"

type Store struct {
	data map[string]string
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

	value, _ := store.Get("a")

	fmt.Println(value)

	fmt.Println("GoKV - Go key value store project")
}
