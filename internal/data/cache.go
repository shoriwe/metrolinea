package data

import "github.com/shoriwe/metrolinea/internal/data/db_objects"

func GenerateCookie(userInformation *db_objects.UserInformation) (string, error) {
	cookies, cookieGenerationError := generateCookieCallback(userInformation)
	if cookieGenerationError != nil {
		go LogCookieGenerationAttempt(userInformation, false)
		return "", cookieGenerationError
	}
	go LogCookieGenerationAttempt(userInformation, true)
	return cookies, nil
}

func CheckCookies(cookies string) (*db_objects.UserInformation, bool, error) {
	userInformation, succeed, checkError := checkCookiesCallback(cookies)
	if checkError != nil {
		go LogCheckCookies(cookies, false)
		return nil, false, checkError
	}
	if succeed {
		go LogCheckCookies(userInformation.Username, true)
		return userInformation, true, nil
	}
	go LogCheckCookies(cookies, false)
	return nil, false, nil
}

func Logout(cookies string) (bool, error) {
	userInformation, validCookies, cookieCheckError := CheckCookies(cookies)
	if cookieCheckError != nil {
		go LogLogoutAttempt(cookies, false)
		return false, cookieCheckError
	}
	if validCookies {
		success, logoutError := logoutCallback(cookies)
		go LogLogoutAttempt(userInformation.Username, success)
		return success, logoutError
	}
	go LogLogoutAttempt(cookies, false)
	return false, nil
}
