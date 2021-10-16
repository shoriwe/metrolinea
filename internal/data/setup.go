package data

import (
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/data/test"
	"net/http"
	"time"
)

type Callbacks interface {
	LogCookieGenerationAttempt(request *http.Request, now time.Time, userInformation *db_objects.UserInformation, succeed bool) error
	LogLoginAttempt(request *http.Request, now time.Time, usernameOrCookies string, succeed bool) error
	LogRegisterAttempt(request *http.Request, now time.Time, username, message string, succeed bool) error
	LogLogoutAttempt(request *http.Request, now time.Time, usernameOrCookies string, succeed bool) error
	LogCheckCookies(request *http.Request, now time.Time, usernameOrCookies string, succeed bool) error
	LogUpdatePasswordAttempt(request *http.Request, now time.Time, usernameOrCookies string, succeed bool) error
	LogUpdateEmailAttempt(request *http.Request, now time.Time, usernameOrCookies string, succeed bool) error
	LogUserExists(request *http.Request, now time.Time, username string, exists bool) error
	LogError(request *http.Request, now time.Time, err error) error
	CheckUserExists(request *http.Request, username string) (bool, error)
	GenerateCookie(request *http.Request, userInformation *db_objects.UserInformation) (string, error)
	CheckCookies(request *http.Request, cookies string) (*db_objects.UserInformation, bool, error)
	Login(request *http.Request, username, password string) (*db_objects.UserInformation, bool, error)
	Register(request *http.Request, registrationForm *forms.RegisterForm) (bool, string, error)
	Logout(request *http.Request, cookies string) (bool, error)
	UpdatePassword(request *http.Request, username string, oldPassword, newPassword string) (bool, string, error)
	UpdateEmail(request *http.Request, username string, password, email string) (bool, string, error)
}

type Controller struct {
	callbacks Callbacks
}

type Settings struct {
}

func Setup(settings Settings) error {
	return nil
}

func TestSetup() *Controller {
	return &Controller{
		callbacks: test.NewCallbacks(),
	}
}