package forms

import "time"

type RegisterForm struct {
	Username  string
	Password  string
	Name      string
	BirthDate time.Time
	Email     string
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
	Id        uint
	Kind      uint
	Username  string
	Name      string
	BirthDate time.Time
	Email     string
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

type UpdateEmailForm struct {
	Cookies  string
	NewEmail string
	Password string
}

type UpdateResponse struct {
	Succeed bool
	Message string
}
