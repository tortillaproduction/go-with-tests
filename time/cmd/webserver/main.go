package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	poker "github.com/quii/learn-go-with-tests/command-line/v3"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		log.Fatalf("problem openig %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	server := poker.NewPlayerServer(store)

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
