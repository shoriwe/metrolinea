package database

import (
	"log"
	"metrolinea/internal/database/db_objects"
	"time"
)

/*
	Log functions should always be used with keyboard "go"
*/
func LogError(err error) {
	logError := logErrorCallback(time.Now(), err)
	if logError != nil {
		log.Println(logError)
	}
}

func LogLoginAttempt(userInformation db_objects.UserInformation, succeeded bool) {
	logError := logLoginAttemptCallback(userInformation, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func DisableCookie(cookie string) error {
	return disableCookieCallback(cookie)
}

func GenerateCookie(userInformation db_objects.UserInformation) (string, error) {
	cookie, cookieGenerationError := generateCookieCallback(userInformation)
	if cookieGenerationError != nil {
		// ToDo: Log that the cookie was not generated for userInformation
		return "", cookieGenerationError
	}
	// ToDo: Log that the cookie was successfully generated for userInformation
	return cookie, nil
}

func CheckCookies(cookies string) (db_objects.UserInformation, bool, error) {
	return checkCookiesCallback(cookies)
}

func Login(username, password string) (string, bool, error) {
	userInformation, loginSuccess, loginError := loginCallback(username, password)
	if loginSuccess && loginError == nil {
		go LogLoginAttempt(userInformation, true)
		cookie, cookieGenerationError := GenerateCookie(userInformation)
		return cookie, true, cookieGenerationError
	}
	go LogLoginAttempt(userInformation, loginSuccess)
	return "", loginSuccess, loginError
}
