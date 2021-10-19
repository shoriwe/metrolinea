package data

import "net/http"

func (controller *Controller) ListTerminals(request *http.Request, cookies string) ([]string, bool, error) {
	userInformation, authSucceed, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogListTerminalsAttempt(request, cookies, false)
		return nil, false, checkError
	}
	if authSucceed {
		go controller.LogListTerminalsAttempt(request, userInformation.Username, false)
		return controller.graph.ListNodes(), true, nil
	}
	go controller.LogListTerminalsAttempt(request, cookies, false)
	return nil, false, nil
}
