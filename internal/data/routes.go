package data

import (
	"github.com/shoriwe/metrolinea/internal/data/graph"
	"net/http"
)

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

func (controller *Controller) ListRoutes(request *http.Request, cookies string) (map[string]graph.Route, bool, error) {
	userInformation, authSucceed, checkError := controller.CheckCookies(request, cookies)
	if checkError != nil {
		go controller.LogListTerminalsAttempt(request, cookies, false)
		return nil, false, checkError
	}
	if authSucceed {
		go controller.LogListTerminalsAttempt(request, userInformation.Username, false)
		return controller.graph.ListRoutes(), true, nil
	}
	go controller.LogListTerminalsAttempt(request, cookies, false)
	return nil, false, nil
}
