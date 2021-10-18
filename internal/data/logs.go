package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"log"
	"net/http"
	"time"
)

/*
	Log functions should always be used with keyboard "go"
*/
func (controller *Controller) LogError(request *http.Request, err error) {
	logError := controller.callbacks.LogError(request, time.Now(), err)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogUserExists(request *http.Request, username string, exists bool) {
	logError := controller.callbacks.LogUserExists(request, time.Now(), username, exists)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogLoginAttempt(request *http.Request, username string, succeeded bool) {
	logError := controller.callbacks.LogLoginAttempt(request, time.Now(), username, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogCheckCookies(request *http.Request, usernameOrCookies string, succeeded bool) {
	logError := controller.callbacks.LogCheckCookies(request, time.Now(), usernameOrCookies, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogLogoutAttempt(request *http.Request, usernameOrCookies string, succeeded bool) {
	logError := controller.callbacks.LogLogoutAttempt(request, time.Now(), usernameOrCookies, succeeded)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogRegistrationAttempt(request *http.Request, username, message string, succeed bool) {
	logError := controller.callbacks.LogRegisterAttempt(request, time.Now(), username, message, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogCookieGenerationAttempt(request *http.Request, userInformation *db_objects.UserInformation, succeed bool) {
	logError := controller.callbacks.LogCookieGenerationAttempt(request, time.Now(), userInformation, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogUpdatePasswordAttempt(request *http.Request, username string, succeed bool) {
	logError := controller.callbacks.LogUpdatePasswordAttempt(request, time.Now(), username, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogUpdateEmailAttempt(request *http.Request, username string, succeed bool) {
	logError := controller.callbacks.LogUpdateEmailAttempt(request, time.Now(), username, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogAdminUpdateUserPasswordAttempt(request *http.Request, usernameOrCookie, targetUsername string, succeed bool) {
	logError := controller.callbacks.LogAdminUpdateUserPasswordAttempt(request, time.Now(), usernameOrCookie, targetUsername, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogAdminUpdateUserEmailAttempt(request *http.Request, usernameOrCookie, targetUsername string, succeed bool) {
	logError := controller.callbacks.LogAdminUpdateUserEmailAttempt(request, time.Now(), usernameOrCookie, targetUsername, succeed)
	if logError != nil {
		log.Println(logError)
	}
}

func (controller *Controller) LogAdminCreateUserAttempt(request *http.Request, usernameOrCookie, targetUsername string, succeed bool) {
	logError := controller.callbacks.LogAdminCreateUserAttempt(request, time.Now(), usernameOrCookie, targetUsername, succeed)
	if logError != nil {
		log.Println(logError)
	}
}
