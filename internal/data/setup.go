package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/data/test"
	"time"
)

var (
	logCookieGenerationAttemptCallback func(now time.Time, userInformation *db_objects.UserInformation, succeed bool) error
	logLoginAttemptCallback            func(now time.Time, usernameOrCookies string, succeed bool) error
	logLogoutAttemptCallback           func(now time.Time, usernameOrCookies string, succeed bool) error
	logWhoamiAttemptCallback           func(now time.Time, usernameOrCookies string, succeed bool) error
	logCheckCookiesCallback            func(now time.Time, usernameOrCookies string, succeed bool) error
	logUserExistsCallback              func(now time.Time, username string, exists bool) error
	logErrorCallback                   func(now time.Time, err error) error
)

var (
	checkUserExistsCallback func(username string) (bool, error)
	generateCookieCallback  func(userInformation *db_objects.UserInformation) (string, error)
	checkCookiesCallback    func(cookies string) (*db_objects.UserInformation, bool, error)
	loginCallback           func(username, password string) (*db_objects.UserInformation, bool, error)
	whoamiCallback          func(userInformation *db_objects.UserInformation) (*db_objects.Whoami, bool, error)
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
	logWhoamiAttemptCallback = test.LogWhoamiAttempt
	logUserExistsCallback = test.LogUserExists
	logCookieGenerationAttemptCallback = test.LogCookieGenerationAttempt
	logCheckCookiesCallback = test.LogCheckCookies

	// Setup general functionality
	checkUserExistsCallback = test.CheckUserExists
	checkCookiesCallback = test.CheckCookies

	// Setup Account functionality
	generateCookieCallback = test.GenerateCookie
	whoamiCallback = test.Whoami
	loginCallback = test.Login
	logoutCallback = test.Logout
}
