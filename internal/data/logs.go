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

func LogUserExists(username string, exists bool) {
	logError := logUserExistsCallback(time.Now(), username, exists)
	if logError != nil {
		log.Println(logError)
	}
}

func LogLoginAttempt(username string, succeeded bool) {
	logError := logLoginAttemptCallback(time.Now(), username, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func LogCheckCookies(usernameOrCookies string, succeeded bool) {
	logError := logCheckCookiesCallback(time.Now(), usernameOrCookies, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func LogLogoutAttempt(usernameOrCookies string, succeeded bool) {
	logError := logLogoutAttemptCallback(time.Now(), usernameOrCookies, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func LogRegistrationAttempt(username, message string, succeed bool) {
	logError := logRegisterAttemptCallback(time.Now(), username, message, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func LogCookieGenerationAttempt(userInformation *db_objects.UserInformation, succeed bool) {
	logError := logCookieGenerationAttemptCallback(time.Now(), userInformation, succeed)
	if logError != nil {
		log.Println(logError)
	}
}
