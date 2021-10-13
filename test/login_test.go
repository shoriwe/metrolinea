package test

import (
	"bytes"
	"encoding/json"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	content_types "github.com/shoriwe/metrolinea/internal/content-types"
	"io"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	loginForm, marshalError := json.Marshal(
		forms.LoginForm{
			Username: "John",
			Password: "Connor",
		},
	)
	if marshalError != nil {
		t.Fatal(marshalError)
	}
	response, postError := server.Client().Post(server.URL+"/login", content_types.Json, bytes.NewReader(loginForm))
	if postError != nil {
		t.Fatal(postError)
	}
	if response.StatusCode != http.StatusOK {
		t.Fatal("Login failed")
	}
	responseBody, readError := io.ReadAll(response.Body)
	if readError != nil {
		t.Fatal(readError)
	}
	var loginResponse forms.LoginResult
	unmarshalError := json.Unmarshal(responseBody, &loginResponse)
	if unmarshalError != nil {
		t.Fatal(unmarshalError)
	}
}
