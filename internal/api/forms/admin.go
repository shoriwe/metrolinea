package forms

import (
	"github.com/shoriwe/metrolinea/internal/data/graph"
	"time"
)

type AdminUpdateUserPasswordForm struct {
	Cookies     string
	Username    string
	NewPassword string
}

type AdminUpdateUserEmailForm struct {
	Cookies  string
	Username string
	NewEmail string
}

type AdminCreateUserForm struct {
	Cookies   string
	Kind      uint
	Username  string
	Password  string
	Name      string
	BirthDate time.Time
	Email     string
}

type AdminDisableUserForm struct {
	Cookies  string
	Username string
}

type AdminAddTerminalsForm struct {
	Cookies   string
	Terminals []string
}

type AdminAddRoutesForm struct {
	Cookies string
	Routes  map[string]graph.Route
}
