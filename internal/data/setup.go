package data

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/data/test"
	"time"
)

var (
	logCookieGenerationAttemptCallback func(time time.Time, userInformation db_objects.UserInformation, succeed bool) error
	logLoginAttemptCallback            func(time time.Time, userInformation db_objects.UserInformation, succeed bool) error
	logLogoutAttemptCallback           func(time time.Time, userInformation db_objects.UserInformation, succeed bool) error
	logWhoamiAttemptCallback           func(time time.Time, whoami db_objects.Whoami, succeed bool) error
	logErrorCallback                   func(time time.Time, err error) error
	disableCookieCallback              func(cookies string) error
	generateCookieCallback             func(userInformation db_objects.UserInformation) (string, error)
	checkCookiesCallback               func(cookies string) (db_objects.UserInformation, bool, error)
	loginCallback                      func(username, password string) (db_objects.UserInformation, bool, error)
	whoamiCallback                     func(cookies string) (db_objects.Whoami, bool, error)
	logoutCallback                     func(cookies string) (db_objects.UserInformation, bool, error)
)

type Settings struct {
}

func Setup(settings Settings) error {
	return nil
}

func TestSetup() {
	logLoginAttemptCallback = test.LogLoginAttempt
	logLogoutAttemptCallback = test.LogLogoutAttempt
	logWhoamiAttemptCallback = test.LogWhoamiAttempt
	logCookieGenerationAttemptCallback = test.LogCookieGenerationAttempt
	generateCookieCallback = test.GenerateCookie
	logErrorCallback = test.LogError
	whoamiCallback = test.Whoami
	loginCallback = test.Login
	logoutCallback = test.Logout
}
