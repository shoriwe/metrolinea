package data

import "github.com/shoriwe/metrolinea/internal/api/forms"

func Login(username, password string) (string, bool, error) {
	userInformation, loginSuccess, loginError := loginCallback(username, password)
	if loginError != nil {
		go LogLoginAttempt(username, false)
		return "", false, loginError
	}
	if loginSuccess {
		go LogLoginAttempt(username, true)
		cookies, cookieGenerationError := GenerateCookie(userInformation)
		return cookies, true, cookieGenerationError
	}
	go LogLoginAttempt(username, false)
	return "", false, nil
}

func Register(registrationForm *forms.RegisterForm) (bool, string, error) {
	doesAlreadyExists, checkError := CheckUserExists(registrationForm.Username)
	if checkError != nil {
		return false, "", checkError
	}
	if doesAlreadyExists {
		return false, "Username already taken", nil
	}
	success, message, registrationError := registerCallback(registrationForm)
	if registrationError != nil {
		go LogRegistrationAttempt(registrationForm.Username, message, false)
		return false, "", registrationError
	}
	if success {
		go LogRegistrationAttempt(registrationForm.Username, "", true)
		return true, message, nil
	}
	go LogRegistrationAttempt(registrationForm.Username, message, false)
	return false, message, nil
}
