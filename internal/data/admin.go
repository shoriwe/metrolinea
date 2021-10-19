package data

import (
	"fmt"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"net/http"
)

func (controller *Controller) AdminUpdateUserPassword(request *http.Request, cookies, username, newPassword string) (bool, string, error) {
	userInformation, validCookies, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogAdminUpdateUserPasswordAttempt(request, cookies, username, false)
		return false, "Something goes wrong!", checkError
	}
	if validCookies {
		if userInformation.Kind != db_objects.Administrator {
			go controller.LogAdminUpdateUserPasswordAttempt(request, cookies, username, false)
			return false, "access denied, functionality just for admins!", nil
		}
		succeed, message, updateError := controller.callbacks.AdminUpdateUserPassword(request, username, newPassword)
		if updateError != nil {
			go controller.LogAdminUpdateUserPasswordAttempt(request, userInformation.Username, username, false)
			return false, "Something goes wrong!", updateError
		}
		if succeed {
			go controller.LogAdminUpdateUserPasswordAttempt(request, userInformation.Username, username, true)
		} else {
			go controller.LogAdminUpdateUserPasswordAttempt(request, userInformation.Username, username, false)
		}
		return succeed, message, nil

	}
	go controller.LogAdminUpdateUserPasswordAttempt(request, cookies, username, false)
	return false, "invalid cookies", nil
}

func (controller *Controller) AdminUpdateUserEmail(request *http.Request, cookies, username, newEmail string) (bool, string, error) {
	userInformation, validCookies, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogAdminUpdateUserEmailAttempt(request, cookies, username, false)
		return false, "Something goes wrong!", checkError
	}
	if validCookies {
		if userInformation.Kind != db_objects.Administrator {
			go controller.LogAdminUpdateUserEmailAttempt(request, cookies, username, false)
			return false, "access denied, functionality just for admins!", nil
		}
		succeed, message, updateError := controller.callbacks.AdminUpdateUserEmail(request, username, newEmail)
		if updateError != nil {
			go controller.LogAdminUpdateUserEmailAttempt(request, userInformation.Username, username, false)
			return false, "Something goes wrong!", updateError
		}
		if succeed {
			go controller.LogAdminUpdateUserEmailAttempt(request, userInformation.Username, username, true)
		} else {
			go controller.LogAdminUpdateUserEmailAttempt(request, userInformation.Username, username, false)
		}
		return succeed, message, nil

	}
	go controller.LogAdminUpdateUserEmailAttempt(request, cookies, username, false)
	return false, "invalid cookies", nil
}

func (controller *Controller) AdminCreateUser(request *http.Request, createUserForm forms.AdminCreateUserForm) (bool, string, error) {
	userInformation, validCookies, checkError := controller.CheckCookies(request, createUserForm.Cookies)
	if checkError != nil {
		go controller.LogAdminCreateUserAttempt(request, createUserForm.Cookies, createUserForm.Username, false)
		return false, "Something goes wrong!", checkError
	}
	if validCookies {
		if userInformation.Kind != db_objects.Administrator {
			go controller.LogAdminCreateUserAttempt(request, userInformation.Username, createUserForm.Username, false)
			return false, "access denied, functionality just for admins!", nil
		}
		succeed, message, updateError := controller.callbacks.AdminCreateUser(request, createUserForm)
		if updateError != nil {
			go controller.LogAdminCreateUserAttempt(request, userInformation.Username, createUserForm.Username, false)
			return false, "Something goes wrong!", updateError
		}
		if succeed {
			go controller.LogAdminCreateUserAttempt(request, userInformation.Username, createUserForm.Username, true)
		} else {
			go controller.LogAdminCreateUserAttempt(request, userInformation.Username, createUserForm.Username, false)
		}
		return succeed, message, nil

	}
	go controller.LogAdminCreateUserAttempt(request, createUserForm.Cookies, createUserForm.Username, false)
	return false, "invalid cookies", nil
}

func (controller *Controller) AdminDisableUser(request *http.Request, cookies, username string) (bool, string, error) {
	userInformation, validCookies, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogAdminDisableUserAttempt(request, cookies, username, false)
		return false, "Something goes wrong!", checkError
	}
	if validCookies {
		if userInformation.Kind != db_objects.Administrator {
			go controller.LogAdminDisableUserAttempt(request, userInformation.Username, username, false)
			return false, "access denied, functionality just for admins!", nil
		}
		succeed, message, updateError := controller.callbacks.AdminDisableUser(request, username)
		if updateError != nil {
			go controller.LogAdminDisableUserAttempt(request, userInformation.Username, username, false)
			return false, "Something goes wrong!", updateError
		}
		if succeed {
			go controller.LogAdminDisableUserAttempt(request, userInformation.Username, username, true)
		} else {
			go controller.LogAdminDisableUserAttempt(request, userInformation.Username, username, false)
		}
		return succeed, message, nil

	}
	go controller.LogAdminDisableUserAttempt(request, cookies, username, false)
	return false, "invalid cookies", nil
}

func (controller *Controller) AdminAddTerminals(request *http.Request, cookies string, terminals []string) (bool, string, error) {
	userInformation, validCookies, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogAdminAddTerminalsAttempt(request, cookies, terminals, false)
		return false, "Something goes wrong!", checkError
	}
	if validCookies {
		if userInformation.Kind != db_objects.Administrator {
			go controller.LogAdminAddTerminalsAttempt(request, userInformation.Username, terminals, false)
			return false, "access denied, functionality just for admins!", nil
		}
		succeed, existingTerminal := controller.graph.AddNodes(terminals)
		var message string
		if succeed {
			message = "successfully added all the terminals"
			go controller.LogAdminAddTerminalsAttempt(request, userInformation.Username, terminals, true)
		} else {
			message = fmt.Sprintf("Terminal with name %s already exists", existingTerminal)
			go controller.LogAdminAddTerminalsAttempt(request, userInformation.Username, terminals, false)
		}
		return succeed, message, nil

	}
	go controller.LogAdminAddTerminalsAttempt(request, cookies, terminals, false)
	return false, "invalid cookies", nil
}
