package test

import (
	"github.com/shoriwe/metrolinea/internal/api"
	"github.com/shoriwe/metrolinea/internal/database"
	"net/http/httptest"
)

var (
	server *httptest.Server
)

func init() {
	database.TestSetup()
	server = httptest.NewServer(api.MetrolineaHandler)
}
