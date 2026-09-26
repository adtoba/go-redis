package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {

}

func (s *Store) Set(key, value string) bool {

}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {
	fmt.Println("GoKV - Go key value store project")
}
