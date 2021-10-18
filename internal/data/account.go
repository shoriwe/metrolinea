package data

import (
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"net/http"
)

func (controller *Controller) Login(request *http.Request, username, password string) (string, bool, error) {
	userInformation, loginSuccess, loginError := controller.callbacks.Login(request, username, password)
	if loginError != nil {
		go controller.LogLoginAttempt(request, username, false)
		return "", false, loginError
	}
	if loginSuccess {
		go controller.LogLoginAttempt(request, username, true)
		cookies, cookieGenerationError := controller.GenerateCookie(request, userInformation)
		return cookies, true, cookieGenerationError
	}
	go controller.LogLoginAttempt(request, username, false)
	return "", false, nil
}

func (controller *Controller) Register(request *http.Request, registrationForm *forms.RegisterForm) (bool, string, error) {
	doesAlreadyExists, checkError := controller.CheckUserExists(request, registrationForm.Username)
	if checkError != nil {
		return false, "", checkError
	}
	if doesAlreadyExists {
		return false, "Username already taken", nil
	}
	success, message, registrationError := controller.callbacks.Register(request, registrationForm)
	if registrationError != nil {
		go controller.LogRegistrationAttempt(request, registrationForm.Username, message, false)
		return false, "", registrationError
	}
	if success {
		go controller.LogRegistrationAttempt(request, registrationForm.Username, "", true)
		return true, message, nil
	}
	go controller.LogRegistrationAttempt(request, registrationForm.Username, message, false)
	return false, message, nil
}

func (controller *Controller) UpdatePassword(request *http.Request, cookies, oldPassword, newPassword string) (bool, string, error) {
	userInformation, succeed, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogUpdatePasswordAttempt(request, cookies, false)
		return false, "", checkError
	}
	if succeed {
		updateSucceed, message, updateError := controller.callbacks.UpdatePassword(request, userInformation.Username, oldPassword, newPassword)
		if updateError != nil {
			go controller.LogUpdatePasswordAttempt(request, userInformation.Username, false)
			return false, "Internal server error", nil
		}
		if updateSucceed {
			go controller.LogUpdatePasswordAttempt(request, userInformation.Username, true)
		} else {
			go controller.LogUpdatePasswordAttempt(request, userInformation.Username, false)
		}
		return updateSucceed, message, nil
	}
	go controller.LogUpdatePasswordAttempt(request, cookies, false)
	return false, "Wrong cookies", nil
}

func (controller *Controller) UpdateEmail(request *http.Request, cookies, password, email string) (bool, string, error) {
	userInformation, succeed, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogUpdateEmailAttempt(request, cookies, false)
		return false, "", checkError
	}
	if succeed {
		updateSucceed, message, updateError := controller.callbacks.UpdateEmail(request, userInformation.Username, password, email)
		if updateError != nil {
			go controller.LogUpdateEmailAttempt(request, userInformation.Username, false)
			return false, "Internal server error", nil
		}
		if updateSucceed {
			go controller.LogUpdateEmailAttempt(request, userInformation.Username, true)
		} else {
			go controller.LogUpdateEmailAttempt(request, userInformation.Username, false)
		}
		return updateSucceed, message, nil
	}
	go controller.LogUpdateEmailAttempt(request, cookies, false)
	return false, "Wrong cookies", nil
}
