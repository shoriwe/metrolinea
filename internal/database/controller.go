package database

import (
	"github.com/shoriwe/metrolinea/internal/database/db_objects"
	"log"
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
	logError := logLoginAttemptCallback(time.Now(), userInformation, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func LogCookieGenerationAttempt(userInformation db_objects.UserInformation, succeed bool) {
	logError := logCookieGenerationAttemptCallback(time.Now(), userInformation, succeed)
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
		go LogCookieGenerationAttempt(userInformation, false)
		return "", cookieGenerationError
	}
	go LogCookieGenerationAttempt(userInformation, true)
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
