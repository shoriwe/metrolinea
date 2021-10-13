package forms

type LoginForm struct {
	Username string
	Password string
}

type LoginResult struct {
	Cookie string
}
