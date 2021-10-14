package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shoriwe/metrolinea/internal/api"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	content_types "github.com/shoriwe/metrolinea/internal/content-types"
	"github.com/shoriwe/metrolinea/internal/data"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	data.TestSetup()
	server := httptest.NewServer(api.MetrolineaHandler)

	var sessionCookies string
	{ // Register Account
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "terminator",
				Password: "Hasta la vista baby!",
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
		var loginResponse forms.LoginResponse
		unmarshalError := json.Unmarshal(responseBody, &loginResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		sessionCookies = loginResponse.Cookies
	}
	{ // Whoami
		whoamiForm, marshalError := json.Marshal(
			forms.WhoamiForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/whoami", content_types.Json, bytes.NewReader(whoamiForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Whoami failed")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var whoami forms.WhoamiResponse
		unmarshalError := json.Unmarshal(responseBody, &whoami)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(whoami)
	}
	{ // Logout
		logoutForm, marshalError := json.Marshal(
			forms.LogoutForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/logout", content_types.Json, bytes.NewReader(logoutForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Logout failed")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var logout forms.LogoutResponse
		unmarshalError := json.Unmarshal(responseBody, &logout)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(logout)
	}
}

func TestRegister(t *testing.T) {
	data.TestSetup()
	server := httptest.NewServer(api.MetrolineaHandler)
	{ // Login
		registrationForm, marshalError := json.Marshal(
			forms.RegisterForm{
				Username:         "lord-protector",
				Password:         "dawnwall1",
				Name:             "Corvo Atano",
				BirthDate:        time.Time{},
				CardNumber:       "0000000",
				Email:            "corvo@akane.fr",
				EmergencyContact: "EMILY_PHONE_NUMBER",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/register", content_types.Json, bytes.NewReader(registrationForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Registration failed")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var registrationResponse forms.RegisterResponse
		unmarshalError := json.Unmarshal(responseBody, &registrationResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		if !registrationResponse.Succeed {
			t.Fatal(registrationResponse.Message)
		}
	}
	var sessionCookies string
	{ // Login
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "lord-protector",
				Password: "dawnwall1",
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
		var loginResponse forms.LoginResponse
		unmarshalError := json.Unmarshal(responseBody, &loginResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		sessionCookies = loginResponse.Cookies
	}
	{ // Whoami
		whoamiForm, marshalError := json.Marshal(
			forms.WhoamiForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/whoami", content_types.Json, bytes.NewReader(whoamiForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Whoami failed")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var whoami forms.WhoamiResponse
		unmarshalError := json.Unmarshal(responseBody, &whoami)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(whoami)
	}
	{ // Logout
		logoutForm, marshalError := json.Marshal(
			forms.LogoutForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/logout", content_types.Json, bytes.NewReader(logoutForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Logout failed")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var logout forms.LogoutResponse
		unmarshalError := json.Unmarshal(responseBody, &logout)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(logout)
	}
}
