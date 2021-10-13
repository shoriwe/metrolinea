package api

import (
	"metrolinea/internal/api/session"
	_ "metrolinea/internal/api/session"
	"net/http"
)

var MetrolineaHandler = &http.ServeMux{}

func init() {
	MetrolineaHandler.HandleFunc("/login", session.Login)
}
