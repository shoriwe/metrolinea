package test

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"log"
	"time"
)

func LogError(_ time.Time, err error) error {
	log.Println(err)
	return nil
}

func LogCheckCookies(_ time.Time, username string, exists bool) error {
	if exists {
		log.Printf("VALID COOKIES FOR USER: %s\n", username)
	} else {
		log.Printf("INVALID COOKIES RECEIVED: %S\n", username)
	}
	return nil
}

func LogUserExists(_ time.Time, username string, exists bool) error {
	if exists {
		log.Printf("USER: %s EXISTS\n", username)
	} else {
		log.Printf("USER: %s DOES NOT EXISTS\n", username)
	}
	return nil
}

func LogLoginAttempt(_ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("LOGIN FOR USER: %s SUCCEED\n", usernameOrCookies)
	} else {
		log.Printf("LOGIN FOR USER: %s FAILED\n", usernameOrCookies)
	}
	return nil
}

func LogLogoutAttempt(_ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("LOGOUT FOR USER: %s SUCCEED\n", usernameOrCookies)
	} else {
		log.Printf("LOGOUT WITH COOKIES: %s FAILED\n", usernameOrCookies)
	}
	return nil
}

func LogWhoamiAttempt(_ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("WHOAMI FOR USER: %s SUCCEED\n", usernameOrCookies)
	} else {
		log.Printf("WHOAMI WITH COOKIES: %s FAILED\n", usernameOrCookies)
	}
	return nil
}

func LogCookieGenerationAttempt(_ time.Time, userInformation *db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Printf("COOKIE GENERATION FOR USER: %s SUCCEED\n", userInformation.Username)
	} else {
		log.Printf("COOKIE GENERATION FOR USER: %s FAILED\n", userInformation.Username)
	}
	return nil
}
