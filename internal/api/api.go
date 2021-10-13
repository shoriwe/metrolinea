package api

import (
	"github.com/shoriwe/metrolinea/internal/api/session"
	_ "github.com/shoriwe/metrolinea/internal/api/session"
	"net/http"
)

var MetrolineaHandler = &http.ServeMux{}

func init() {
	MetrolineaHandler.HandleFunc("/login", session.Login)
	MetrolineaHandler.HandleFunc("/logout", session.Logout)
	MetrolineaHandler.HandleFunc("/whoami", session.Whoami)
}
