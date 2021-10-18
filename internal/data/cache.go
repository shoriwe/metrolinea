package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"net/http"
)

func (controller *Controller) GenerateCookie(request *http.Request, userInformation *db_objects.UserInformation) (string, error) {
	cookies, cookieGenerationError := controller.callbacks.GenerateCookie(request, userInformation)
	if cookieGenerationError != nil {
		go controller.LogCookieGenerationAttempt(request, userInformation, false)
		return "", cookieGenerationError
	}
	go controller.LogCookieGenerationAttempt(request, userInformation, true)
	return cookies, nil
}

func (controller *Controller) CheckCookies(request *http.Request, cookies string) (*db_objects.UserInformation, bool, error) {
	userInformation, succeed, checkError := controller.callbacks.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogCheckCookies(request, cookies, false)
		return nil, false, checkError
	}
	if succeed {
		go controller.LogCheckCookies(request, userInformation.Username, true)
		return userInformation, true, nil
	}
	go controller.LogCheckCookies(request, cookies, false)
	return nil, false, nil
}

func (controller *Controller) Logout(request *http.Request, cookies string) (bool, error) {
	userInformation, validCookies, cookieCheckError := controller.CheckCookies(request, cookies)
	if cookieCheckError != nil {
		go controller.LogLogoutAttempt(request, cookies, false)
		return false, cookieCheckError
	}
	if validCookies {
		success, logoutError := controller.callbacks.Logout(request, cookies)
		go controller.LogLogoutAttempt(request, userInformation.Username, success)
		return success, logoutError
	}
	go controller.LogLogoutAttempt(request, cookies, false)
	return false, nil
}
