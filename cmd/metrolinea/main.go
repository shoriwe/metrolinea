package main

import (
	"log"
	"metrolinea/internal/api"
	"metrolinea/internal/database"
	"net/http"
)

func main() {
	setupError := database.Setup(database.Settings{})
	if setupError != nil {
		log.Fatal(setupError)
	}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.MetrolineaHandler)) // FixMe: Make the ip configurable
}
