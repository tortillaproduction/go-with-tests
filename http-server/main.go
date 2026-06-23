package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", server))
}
