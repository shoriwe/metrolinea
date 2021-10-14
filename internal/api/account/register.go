package account

import (
	"encoding/json"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	content_types "github.com/shoriwe/metrolinea/internal/content-types"
	"github.com/shoriwe/metrolinea/internal/data"
	"github.com/shoriwe/metrolinea/internal/errors"
	"github.com/shoriwe/metrolinea/internal/messages"
	"io"
	"net/http"
)

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		go data.LogError(errors.MethodNotAllowed(request.RemoteAddr, request.Method, request.RequestURI))
		_, writeError := responseWriter.Write(messages.MethodsAllowed(http.MethodPost))
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Handle wrong content type
	if request.Header.Get("content-type") != content_types.Json {
		go data.LogError(errors.ContentTypeNotSupported(request.RemoteAddr, request.Method, request.RequestURI, request.Header.Get("content-type")))
		responseWriter.WriteHeader(http.StatusUnsupportedMediaType)
		_, writeError := responseWriter.Write(messages.ContentTypesSupported(content_types.Json))
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
	body, readError := io.ReadAll(request.Body)
	if readError != nil {
		go data.LogError(errors.GoRuntimeError(readError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
	var registerForm forms.RegisterForm
	unmarshalError := json.Unmarshal(body, &registerForm)
	if unmarshalError != nil {
		go data.LogError(errors.GoRuntimeError(unmarshalError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
	registrationSucceed, registrationMessage, registrationError := data.Register(&registerForm)
	if registrationError != nil {
		go data.LogError(errors.GoRuntimeError(registrationError, request.RemoteAddr, request.Method, request.RequestURI))
		responseWriter.WriteHeader(http.StatusInternalServerError)
		_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
		if writeError != nil {
			go data.LogError(writeError)
			return
		}
		return
	}
	response, responseMarshalError := json.Marshal(
		forms.RegisterResponse{
			Succeed: registrationSucceed,
			Message: registrationMessage,
		},
	)
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
	if !registrationSucceed {
		responseWriter.WriteHeader(http.StatusNotAcceptable)
	}
	responseWriter.Header().Set("content-type", content_types.Json)
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
