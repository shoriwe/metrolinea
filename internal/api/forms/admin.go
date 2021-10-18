package forms

import "time"

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
