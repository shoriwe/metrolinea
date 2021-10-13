package test

import (
	"metrolinea/internal/api"
	"metrolinea/internal/database"
	"net/http/httptest"
)

var (
	server *httptest.Server
)

func init() {
	database.TestSetup()
	server = httptest.NewServer(api.MetrolineaHandler)
}
