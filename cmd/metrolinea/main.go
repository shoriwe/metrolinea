package main

import (
	"github.com/shoriwe/metrolinea/internal/api"
	"github.com/shoriwe/metrolinea/internal/data"
	"log"
	"net/http"
)

func main() {
	controller, setupError := data.Setup(data.Settings{})
	if setupError != nil {
		log.Fatal(setupError)
	}
	// FixMe: make the database configurable
	// FixMe: make the cache configurable
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", api.NewHandler(controller))) // FixMe: Make the ip configurable
}
