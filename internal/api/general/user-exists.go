package general

import (
	"encoding/json"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data"
	"github.com/shoriwe/metrolinea/internal/errors"
	"github.com/shoriwe/metrolinea/internal/messages"
	"net/http"
	"strings"
)

func UserExists(controller *data.Controller) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		argument := strings.Split(strings.TrimPrefix(request.RequestURI, "/user/exists/"), "/")
		if len(argument) != 1 {
			responseWriter.WriteHeader(http.StatusNotFound)
			// ToDo: Log this
			return
		}
		username := argument[0]
		exists, checkError := controller.CheckUserExists(request, username)
		if checkError != nil {
			go controller.LogError(request, errors.GoRuntimeError(checkError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		if !exists {
			responseWriter.WriteHeader(http.StatusNotFound)
			// ToDo: Log this
			return
		}
		response, responseMarshalError := json.Marshal(forms.UserExistsResponse{Exists: true})
		if responseMarshalError != nil {
			go controller.LogError(request, errors.GoRuntimeError(responseMarshalError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		_, responseWriteError := responseWriter.Write(response)
		if responseWriteError != nil {
			go controller.LogError(request, errors.GoRuntimeError(responseWriteError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
	}
}
