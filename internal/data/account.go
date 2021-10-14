package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
)

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

func Whoami(cookies string) (*db_objects.Whoami, bool, error) {
	userInformation, validCookies, cookieCheckError := CheckCookies(cookies)
	if cookieCheckError != nil {
		go LogWhoamiAttempt(cookies, false)
		return nil, false, cookieCheckError
	}
	if validCookies {
		result, success, whoamiError := whoamiCallback(userInformation)
		if whoamiError != nil {
			return nil, false, whoamiError
		}
		go LogWhoamiAttempt(userInformation.Username, success)
		return result, success, nil
	}
	go LogWhoamiAttempt(cookies, false)
	return nil, false, nil
}
