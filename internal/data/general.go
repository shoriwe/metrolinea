package data

import "net/http"

func (controller *Controller) CheckUserExists(request *http.Request, username string) (bool, error) {
	exists, checkError := controller.callbacks.CheckUserExists(request, username)
	go controller.LogUserExists(request, username, exists)
	return exists, checkError
}
