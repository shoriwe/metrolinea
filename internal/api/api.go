package api

import (
	"github.com/shoriwe/metrolinea/internal/api/account"
	_ "github.com/shoriwe/metrolinea/internal/api/account"
	admin "github.com/shoriwe/metrolinea/internal/api/admin"
	"github.com/shoriwe/metrolinea/internal/api/general"
	"github.com/shoriwe/metrolinea/internal/api/routes"
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
	handler.HandleFunc("/admin/update/user/password", admin.UpdateUserPassword(controller))
	handler.HandleFunc("/admin/update/user/email", admin.UpdateUserEmail(controller))
	handler.HandleFunc("/admin/create/user", admin.CreateUser(controller))
	handler.HandleFunc("/admin/disable/user", admin.DisableUser(controller))
	handler.HandleFunc("/admin/add/terminals", admin.AddTerminals(controller))
	handler.HandleFunc("/admin/add/routes", admin.AddRoutes(controller))
	// handler.HandleFunc("/admin/delete/terminals", admin.DeleteTerminals(controller))
	// handler.HandleFunc("/admin/delete/routes", admin.DeleteRoutes(controller))
	handler.HandleFunc("/list/terminals", routes.ListTerminals(controller))
	handler.HandleFunc("/list/routes", routes.ListRoutes(controller))
	// ToDo: handler.HandleFunc("/find/route", routes.FindRoute(controller))
	return handler
}
