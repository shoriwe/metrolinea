package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/shoriwe/metrolinea/internal/api"
	"github.com/shoriwe/metrolinea/internal/api/forms"
	content_types "github.com/shoriwe/metrolinea/internal/content-types"
	"github.com/shoriwe/metrolinea/internal/data"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/data/graph"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAdminUpdateUserPassword(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Update Marla's password
		updatePasswordForm, marshalError := json.Marshal(
			forms.AdminUpdateUserPasswordForm{
				Cookies:     sessionCookies,
				Username:    "mSinger",
				NewPassword: "Hola Mundo!",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/update/user/password", content_types.Json, bytes.NewReader(updatePasswordForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Update password failed")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // Login as marla
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "mSinger",
				Password: "Hola Mundo!",
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
		fmt.Println(string(responseBody))
	}
}

func TestAdminUpdateUserEmail(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Update Marla's password
		updateEmailForm, marshalError := json.Marshal(
			forms.AdminUpdateUserEmailForm{
				Cookies:  sessionCookies,
				Username: "mSinger",
				NewEmail: "friday@day.com",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/update/user/email", content_types.Json, bytes.NewReader(updateEmailForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Update email failed")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // Login as marla
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "mSinger",
				Password: "The first rule of the fight club is...",
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
		fmt.Println(string(responseBody))
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
		if whoami.Email != "friday@day.com" {
			t.Fatal(whoami.Email)
		}
		fmt.Println(string(responseBody))
	}
}

func TestAdminCreateUser(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Create New user
		updateEmailForm, marshalError := json.Marshal(
			forms.AdminCreateUserForm{
				Cookies:   sessionCookies,
				Kind:      db_objects.Administrator,
				Username:  "sulcud",
				Password:  "password",
				Name:      "sulcud",
				BirthDate: time.Time{},
				Email:     "sulcud@my-domain.com",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/create/user", content_types.Json, bytes.NewReader(updateEmailForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("User creation error")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // Login as sulcud
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "sulcud",
				Password: "password",
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
		fmt.Println(string(responseBody))
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
		if whoami.Email != "sulcud@my-domain.com" {
			t.Fatal(whoami.Email)
		}
		fmt.Println(string(responseBody))
	}
}

func TestAdminDisableUser(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Disable Marla's user
		updateEmailForm, marshalError := json.Marshal(
			forms.AdminDisableUserForm{
				Cookies:  sessionCookies,
				Username: "mSinger",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/disable/user", content_types.Json, bytes.NewReader(updateEmailForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("User creation error")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // Fail login as Marla's user
		loginForm, marshalError := json.Marshal(
			forms.LoginForm{
				Username: "mSinger",
				Password: "The first rule of the fight club is...",
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/login", content_types.Json, bytes.NewReader(loginForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode == http.StatusOK {
			t.Fatal("Disable failed")
		}
	}
}

func TestAddTerminals(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Add terminals
		addTerminalsForm, marshalError := json.Marshal(
			forms.AdminAddTerminalsForm{
				Cookies:   sessionCookies,
				Terminals: []string{"A", "B", "C"},
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/add/terminals", content_types.Json, bytes.NewReader(addTerminalsForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Add terminals error")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // List terminals
		loginForm, marshalError := json.Marshal(
			forms.ListForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/list/terminals", content_types.Json, bytes.NewReader(loginForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Terminals Add error")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var listTerminalsReponse forms.ListTerminalsResponse
		unmarshalError := json.Unmarshal(responseBody, &listTerminalsReponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		if len(listTerminalsReponse.Terminals) != 3 {
			t.Fatal(string(responseBody))
		}
		aFound := false
		bFound := false
		cFound := false
		for _, element := range listTerminalsReponse.Terminals {
			switch element {
			case "A":
				aFound = true
			case "B":
				bFound = true
			case "C":
				cFound = true
			}
		}
		if !(aFound && bFound && cFound) {
			t.Fatal(listTerminalsReponse.Terminals)
		}
	}
}

func TestAddRoutes(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))

	var sessionCookies string
	{ // Login
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
	{ // Add terminals
		addTerminalsForm, marshalError := json.Marshal(
			forms.AdminAddTerminalsForm{
				Cookies:   sessionCookies,
				Terminals: []string{"A", "B", "C"},
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/add/terminals", content_types.Json, bytes.NewReader(addTerminalsForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Add terminals error")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // Add routes
		addRoutesForm, marshalError := json.Marshal(
			forms.AdminAddRoutesForm{
				Cookies: sessionCookies,
				Routes: map[string]graph.Route{
					"1": {
						Source: "A",
						Length: 1,
						Target: "B",
					},
					"2": {
						Source: "B",
						Length: 1,
						Target: "A",
					},
					"3": {
						Source: "C",
						Length: 5,
						Target: "A",
					},
				},
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/admin/add/routes", content_types.Json, bytes.NewReader(addRoutesForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Add routes error")
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
		if !updateResponse.Succeed {
			t.Fatal(updateResponse.Message)
		}
		fmt.Println(string(responseBody))
	}
	{ // List terminals
		loginForm, marshalError := json.Marshal(
			forms.ListForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/list/terminals", content_types.Json, bytes.NewReader(loginForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Terminals Add error")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var listTerminalsReponse forms.ListTerminalsResponse
		unmarshalError := json.Unmarshal(responseBody, &listTerminalsReponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		if len(listTerminalsReponse.Terminals) != 3 {
			t.Fatal(string(responseBody))
		}
		aFound := false
		bFound := false
		cFound := false
		for _, element := range listTerminalsReponse.Terminals {
			switch element {
			case "A":
				aFound = true
			case "B":
				bFound = true
			case "C":
				cFound = true
			}
		}
		if !(aFound && bFound && cFound) {
			t.Fatal(listTerminalsReponse.Terminals)
		}
	}
	{ // List routes
		loginForm, marshalError := json.Marshal(
			forms.ListForm{
				Cookies: sessionCookies,
			},
		)
		if marshalError != nil {
			t.Fatal(marshalError)
		}
		response, postError := server.Client().Post(server.URL+"/list/routes", content_types.Json, bytes.NewReader(loginForm))
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("Routes Add error")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var listRoutesResponse forms.ListRoutesResponse
		unmarshalError := json.Unmarshal(responseBody, &listRoutesResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		if len(listRoutesResponse.Routes) != 3 {
			t.Fatal(string(responseBody))
		}
	}
}
