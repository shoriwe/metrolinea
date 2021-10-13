package main

import (
	"github.com/shoriwe/metrolinea/internal/api"
	"github.com/shoriwe/metrolinea/internal/database"
	"log"
	"net/http"
)

func main() {
	setupError := database.Setup(database.Settings{})
	if setupError != nil {
		log.Fatal(setupError)
	}
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.MetrolineaHandler)) // FixMe: Make the ip configurable
}
