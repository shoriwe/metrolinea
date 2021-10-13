package forms

import "time"

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
	Number    string
}

type LogoutForm struct {
	Cookies string
}

type LogoutResponse struct {
	Succeed bool
}
