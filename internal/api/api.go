package api

import (
	"github.com/shoriwe/metrolinea/internal/api/account"
	_ "github.com/shoriwe/metrolinea/internal/api/account"
	"github.com/shoriwe/metrolinea/internal/api/general"
	"net/http"
)

var MetrolineaHandler = &http.ServeMux{}

func init() {
	MetrolineaHandler.HandleFunc("/login", account.Login)
	MetrolineaHandler.HandleFunc("/logout", account.Logout)
	MetrolineaHandler.HandleFunc("/whoami", account.Whoami)
	MetrolineaHandler.HandleFunc("/user/exists/", general.UserExists)
}
