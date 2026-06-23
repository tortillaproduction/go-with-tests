package main

import (
	"fmt"
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func (i *InMemoryPlayerStore) RecordWin(name string) {}

func main() {
	server := &PlayerServer{&InMemoryPlayerStore{}}

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
