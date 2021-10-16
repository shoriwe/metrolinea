package api

import (
	"github.com/shoriwe/metrolinea/internal/api/account"
	_ "github.com/shoriwe/metrolinea/internal/api/account"
	"github.com/shoriwe/metrolinea/internal/api/general"
	"github.com/shoriwe/metrolinea/internal/data"
	"net/http"
)

func NewHandler(controller *data.Controller) *http.ServeMux {
	handler := &http.ServeMux{}
	handler.HandleFunc("/login", account.Login(controller))
	handler.HandleFunc("/logout", account.Logout(controller))
	handler.HandleFunc("/whoami", account.Whoami(controller))
	handler.HandleFunc("/register", account.Register(controller))
	handler.HandleFunc("/user/update/password", account.UpdatePassword(controller))
	handler.HandleFunc("/user/update/email", account.UpdateEmail(controller))
	handler.HandleFunc("/user/exists/", general.UserExists(controller))
	return handler
}
