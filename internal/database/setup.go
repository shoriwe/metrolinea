package database

import (
	"github.com/shoriwe/metrolinea/internal/database/db_objects"
	"github.com/shoriwe/metrolinea/internal/database/test"
	"time"
)

var (
	logLoginAttemptCallback func(userInformation db_objects.UserInformation, succeed bool) error
	logErrorCallback        func(time time.Time, err error) error
	disableCookieCallback   func(cookie string) error
	generateCookieCallback  func(userInformation db_objects.UserInformation) (string, error)
	checkCookiesCallback    func(cookie string) (db_objects.UserInformation, bool, error)
	loginCallback           func(username, password string) (db_objects.UserInformation, bool, error)
)

type Settings struct {
}

func Setup(settings Settings) error {
	return nil
}

func TestSetup() {
	logLoginAttemptCallback = test.LogLoginAttempt
	generateCookieCallback = test.GenerateCookie
	logErrorCallback = test.LogError
	loginCallback = test.Login
}
