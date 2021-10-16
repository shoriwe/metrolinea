package forms

import "time"

type RegisterForm struct {
	Username         string
	Password         string
	Name             string
	BirthDate        time.Time
	CardNumber       string
	Email            string
	EmergencyContact string
}

type RegisterResponse struct {
	Succeed bool
	Message string
}

type LoginForm struct {
	Username string
	Password string
}

type LoginResponse struct {
	Cookies string
}

type WhoamiForm struct {
	Cookies string
}

type WhoamiResponse struct {
	Id               uint
	Kind             uint
	Username         string
	Name             string
	BirthDate        time.Time
	Number           string
	Email            string
	EmergencyContact string
}

type LogoutForm struct {
	Cookies string
}

type LogoutResponse struct {
	Succeed bool
}

type UpdatePasswordForm struct {
	Cookies     string
	OldPassword string
	NewPassword string
}

type UpdateEmail struct {
	Cookies  string
	NewEmail string
	Password string
}

type UpdateEmergencyContact struct {
	Cookies             string
	NewEmergencyContact string
	Password            string
}

type UpdateCard struct {
	Cookies string
	NewCard string
}

type UpdateResponse struct {
	Succeed bool
	Message string
}
