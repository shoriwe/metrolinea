package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
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

func LogLogoutAttempt(userInformation db_objects.UserInformation, succeeded bool) {
	logError := logLogoutAttemptCallback(time.Now(), userInformation, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func LogWhoamiAttempt(whoami db_objects.Whoami, succeed bool) {
	logError := logWhoamiAttemptCallback(time.Now(), whoami, succeed)
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

func DisableCookie(cookies string) error {
	return disableCookieCallback(cookies)
}

func GenerateCookie(userInformation db_objects.UserInformation) (string, error) {
	cookies, cookieGenerationError := generateCookieCallback(userInformation)
	if cookieGenerationError != nil {
		go LogCookieGenerationAttempt(userInformation, false)
		return "", cookieGenerationError
	}
	go LogCookieGenerationAttempt(userInformation, true)
	return cookies, nil
}

func CheckCookies(cookies string) (db_objects.UserInformation, bool, error) {
	return checkCookiesCallback(cookies)
}

func Login(username, password string) (string, bool, error) {
	userInformation, loginSuccess, loginError := loginCallback(username, password)
	if loginSuccess && loginError == nil {
		go LogLoginAttempt(userInformation, true)
		cookies, cookieGenerationError := GenerateCookie(userInformation)
		return cookies, true, cookieGenerationError
	}
	go LogLoginAttempt(userInformation, loginSuccess)
	return "", loginSuccess, loginError
}

func Logout(cookies string) (bool, error) {
	userInformation, success, logoutError := logoutCallback(cookies)
	if success && logoutError == nil {
		go LogLogoutAttempt(userInformation, true)
		return true, nil
	}
	go LogLogoutAttempt(userInformation, success)
	return success, logoutError
}

func Whoami(cookies string) (db_objects.Whoami, bool, error) {
	whoami, success, logoutError := whoamiCallback(cookies)
	if success && logoutError == nil {
		go LogWhoamiAttempt(whoami, true)
		return whoami, true, nil
	}
	go LogWhoamiAttempt(whoami, success)
	return whoami, success, logoutError
}
