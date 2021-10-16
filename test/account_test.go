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
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

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
		fmt.Println(string(responseBody))
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
		fmt.Println(string(responseBody))
	}
}

func TestRegister(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))
	{ // Register
		registrationForm, marshalError := json.Marshal(
			forms.RegisterForm{
				Username:  "lord-protector",
				Password:  "dawnwall1",
				Name:      "Corvo Atano",
				BirthDate: time.Time{},
				Email:     "corvo@akane.fr",
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
		fmt.Println(string(responseBody))
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
		fmt.Println(string(responseBody))
	}
}

func TestUpdatePassword(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))
	{ // Register
		registrationForm, marshalError := json.Marshal(
			forms.RegisterForm{
				Username:  "emily",
				Password:  "dawnwall2",
				Name:      "Emily Colwin",
				BirthDate: time.Time{},
				Email:     "emily@akane.fr",
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
				Username: "emily",
				Password: "dawnwall2",
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
	{ // Update Password
		updatePasswordForm, marshalError := json.Marshal(
			forms.UpdatePasswordForm{
				Cookies:     sessionCookies,
				OldPassword: "dawnwall2",
				NewPassword: "my-new-password",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/user/update/password", content_types.Json, bytes.NewReader(updatePasswordForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Update password error")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var updateResponse forms.UpdateResponse
		unmarshalError := json.Unmarshal(responseBody, &updateResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(string(responseBody))
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
		fmt.Println(string(responseBody))
	}
	{ // Login again
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "emily",
				Password: "my-new-password",
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
}

func TestUpdateEmail(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))
	{ // Register
		registrationForm, marshalError := json.Marshal(
			forms.RegisterForm{
				Username:  "emily",
				Password:  "dawnwall2",
				Name:      "Emily Colwin",
				BirthDate: time.Time{},
				Email:     "emily@akane.fr",
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
				Username: "emily",
				Password: "dawnwall2",
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
	{ // Update Email
		updateEmailForm, marshalError := json.Marshal(
			forms.UpdateEmailForm{
				Cookies:  sessionCookies,
				NewEmail: "corp@org.corp",
				Password: "dawnwall2",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/user/update/email", content_types.Json, bytes.NewReader(updateEmailForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Update email error")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var updateResponse forms.UpdateResponse
		unmarshalError := json.Unmarshal(responseBody, &updateResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(string(responseBody))
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
		if whoami.Email != "corp@org.corp" {
			t.Fatal(whoami.Email)
		}
		fmt.Println(string(responseBody))
	}
}
