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

func Logout(controller *data.Controller) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			go controller.LogError(request, errors.MethodNotAllowed(request.RemoteAddr, request.Method, request.RequestURI))
			_, writeError := responseWriter.Write(messages.MethodsAllowed(http.MethodPost))
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			responseWriter.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		// Handle wrong content type
		if request.Header.Get("content-type") != content_types.Json {
			go controller.LogError(request, errors.ContentTypeNotSupported(request.RemoteAddr, request.Method, request.RequestURI, request.Header.Get("content-type")))
			responseWriter.WriteHeader(http.StatusUnsupportedMediaType)
			_, writeError := responseWriter.Write(messages.ContentTypesSupported(content_types.Json))
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		body, readError := io.ReadAll(request.Body)
		if readError != nil {
			go controller.LogError(request, errors.GoRuntimeError(readError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		var logoutForm forms.LogoutForm
		unmarshalError := json.Unmarshal(body, &logoutForm)
		if unmarshalError != nil {
			go controller.LogError(request, errors.GoRuntimeError(unmarshalError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		logoutSuccess, logoutError := controller.Logout(request, logoutForm.Cookies)
		if logoutError != nil {
			go controller.LogError(request, errors.GoRuntimeError(logoutError, request.RemoteAddr, request.Method, request.RequestURI))
			responseWriter.WriteHeader(http.StatusInternalServerError)
			_, writeError := responseWriter.Write(messages.SomethingGoesWrong())
			if writeError != nil {
				go controller.LogError(request, writeError)
				return
			}
			return
		}
		if !logoutSuccess {
			responseWriter.WriteHeader(http.StatusForbidden)
			return
		}
		response, responseMarshalError := json.Marshal(
			forms.LogoutResponse{
				Succeed: true,
			},
		)
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
		responseWriter.Header().Set("content-type", content_types.Json)
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
