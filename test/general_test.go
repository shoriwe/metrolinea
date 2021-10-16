package test

import (
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
)

func TestUserExists(t *testing.T) {
	server := httptest.NewServer(api.NewHandler(data.TestSetup()))
	{ // Check if user exists
		response, postError := server.Client().Post(server.URL+"/user/exists/terminator", content_types.PlainText, nil)
		if postError != nil {
			t.Fatal(postError)
		}
		if response.StatusCode != http.StatusOK {
			t.Fatal("User does not exists")
		}
		responseBody, readError := io.ReadAll(response.Body)
		if readError != nil {
			t.Fatal(readError)
		}
		var existsResponse forms.UserExistsResponse
		unmarshalError := json.Unmarshal(responseBody, &existsResponse)
		if unmarshalError != nil {
			t.Fatal(unmarshalError)
		}
		fmt.Println(string(responseBody))
	}
}
