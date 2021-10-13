package session

import (
	"encoding/json"
	"io"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	content_types "github.com/shoriwe/metrolinea/internal/content-types"
	"github.com/shoriwe/metrolinea/internal/database"
	"github.com/shoriwe/metrolinea/internal/errors"
	"github.com/shoriwe/metrolinea/internal/messages"
	"net/http"
)

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		go database.LogError(errors.MethodNotAllowed(request.RemoteAddr, request.Method, request.RequestURI))
		_, writeError := responseWriter.Write(messages.MethodsAllowed(http.MethodPost))
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Handle wrong content type
	if request.Header.Get("content-type") != content_types.Json {
		go database.LogError(errors.ContentTypeNotSupported(request.RemoteAddr, request.Method, request.RequestURI, request.Header.Get("content-type")))
		responseWriter.WriteHeader(http.StatusUnsupportedMediaType)
		_, writeError := responseWriter.Write(messages.ContentTypesSupported(content_types.Json))
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
	body, readError := io.ReadAll(request.Body)
	if readError != nil {
		go database.LogError(errors.GoRuntimeError(readError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
	var form forms.LoginForm
	unmarshalError := json.Unmarshal(body, &form)
	if unmarshalError != nil {
		go database.LogError(errors.GoRuntimeError(unmarshalError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
	cookie, loginSuccess, loginError := database.Login(form.Username, form.Password)
	if loginError != nil {
		go database.LogError(errors.GoRuntimeError(loginError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
	if !loginSuccess {
		responseWriter.WriteHeader(http.StatusForbidden)
		return
	}
	response, responseMarshalError := json.Marshal(
		forms.LoginResult{
			Cookie: cookie,
		},
	)
	if responseMarshalError != nil {
		go database.LogError(errors.GoRuntimeError(responseMarshalError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
	responseWriter.Header().Set("content-type", content_types.Json)
	_, responseWriteError := responseWriter.Write(response)
	if responseWriteError != nil {
		go database.LogError(errors.GoRuntimeError(responseWriteError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go database.LogError(writeError)
			return
		}
		return
	}
}
