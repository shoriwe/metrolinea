package data

import (
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/data/test"
	"time"
)

var (
	logCookieGenerationAttemptCallback func(now time.Time, userInformation *db_objects.UserInformation, succeed bool) error
	logLoginAttemptCallback            func(now time.Time, usernameOrCookies string, succeed bool) error
	logRegisterAttemptCallback         func(now time.Time, username, message string, succeed bool) error
	logLogoutAttemptCallback           func(now time.Time, usernameOrCookies string, succeed bool) error
	logCheckCookiesCallback            func(now time.Time, usernameOrCookies string, succeed bool) error
	logUserExistsCallback              func(now time.Time, username string, exists bool) error
	logErrorCallback                   func(now time.Time, err error) error
)

var (
	checkUserExistsCallback func(username string) (bool, error)
	generateCookieCallback  func(userInformation *db_objects.UserInformation) (string, error)
	checkCookiesCallback    func(cookies string) (*db_objects.UserInformation, bool, error)
	loginCallback           func(username, password string) (*db_objects.UserInformation, bool, error)
	registerCallback        func(registrationForm *forms.RegisterForm) (bool, string, error)
	logoutCallback          func(cookies string) (bool, error)
)

type Settings struct {
}

func Setup(settings Settings) error {
	return nil
}

func TestSetup() {
	// Setup logs
	logErrorCallback = test.LogError
	logLoginAttemptCallback = test.LogLoginAttempt
	logLogoutAttemptCallback = test.LogLogoutAttempt
	logRegisterAttemptCallback = test.LogRegisterAttempt
	logUserExistsCallback = test.LogUserExists
	logCookieGenerationAttemptCallback = test.LogCookieGenerationAttempt
	logCheckCookiesCallback = test.LogCheckCookies

	// Setup general functionality
	checkUserExistsCallback = test.CheckUserExists
	checkCookiesCallback = test.CheckCookies

	// Setup Account functionality
	generateCookieCallback = test.GenerateCookie
	loginCallback = test.Login
	registerCallback = test.Register
	logoutCallback = test.Logout
}
