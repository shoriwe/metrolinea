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

func UserExists(responseWriter http.ResponseWriter, request *http.Request) {
	argument := strings.Split(strings.TrimPrefix(request.RequestURI, "/user/exists/"), "/")
	if len(argument) != 1 {
		responseWriter.WriteHeader(http.StatusNotFound)
		// ToDo: Log this
		return
	}
	username := argument[0]
	exists, checkError := data.CheckUserExists(username)
	if checkError != nil {
		go data.LogError(errors.GoRuntimeError(checkError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
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
		go data.LogError(errors.GoRuntimeError(responseMarshalError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
	_, responseWriteError := responseWriter.Write(response)
	if responseWriteError != nil {
		go data.LogError(errors.GoRuntimeError(responseWriteError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
}
