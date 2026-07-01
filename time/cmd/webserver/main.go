package main

import (
	"log"
	"net/http"

	poker "github.com/quii/learn-go-with-tests/command-line/v3"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := poker.NewPlayerServer(store)

	log.Println("Starting the server on port 5000")
	log.Println("You can visit http://localhost:5000/play to play a game")

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
