package test

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"log"
	"net/http"
	"time"
)

func (c *Callbacks) LogError(request *http.Request, _ time.Time, err error) error {
	log.Println(request.RemoteAddr, err)
	return nil
}

func (c *Callbacks) LogCheckCookies(request *http.Request, _ time.Time, username string, exists bool) error {
	if exists {
		log.Printf("VALID COOKIES FOR USER: %s - %s\n", username, request.RemoteAddr)
	} else {
		log.Printf("INVALID COOKIES RECEIVED: %s - %s\n", username, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogRegisterAttempt(request *http.Request, _ time.Time, username, message string, succeed bool) error {
	if succeed {
		log.Printf("REGISTRATION FOR USER: %s -%s SUCCEED\n", username, request.RemoteAddr)
	} else {
		log.Printf("REGISTRATION FOR USER: %s - %s FAILED -> %s\n", username, request.RemoteAddr, message)
	}
	return nil
}

func (c *Callbacks) LogUserExists(request *http.Request, _ time.Time, username string, exists bool) error {
	if exists {
		log.Printf("LOGIN FOR USER: %s - %s SUCCEED\n", username, request.RemoteAddr)
	} else {
		log.Printf("LOGIN FOR USER: %s - %s FAILED\n", username, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogLoginAttempt(request *http.Request, _ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("LOGIN FOR USER: %s - %s SUCCEED\n", usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("LOGIN FOR USER: %s - %s FAILED\n", usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogLogoutAttempt(request *http.Request, _ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("LOGOUT FOR USER: %s - %s SUCCEED\n", usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("LOGOUT WITH COOKIES: %s - %s FAILED\n", usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogCookieGenerationAttempt(request *http.Request, _ time.Time, userInformation *db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Printf("COOKIE GENERATION FOR USER: %s - %s SUCCEED\n", userInformation.Username, request.RemoteAddr)
	} else {
		log.Printf("COOKIE GENERATION FOR USER: %s - %s FAILED\n", userInformation.Username, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogUpdatePasswordAttempt(request *http.Request, _ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("UPDATE PASSWORD FOR: %s - %s SUCCEED\n", usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("UPDATE PASSWORD FOR: %s - %s FAILED\n", usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogUpdateEmailAttempt(request *http.Request, _ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("UPDATE EMAIL FOR: %s - %s SUCCEED\n", usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("UPDATE EMAIL FOR: %s - %s FAILED\n", usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminUpdateUserPasswordAttempt(request *http.Request, _ time.Time, usernameOrCookies, targetUsername string, succeed bool) error {
	if succeed {
		log.Printf("UPDATE PASSWORD FOR USERNAME: %s BY ADMIN USER %s - %s SUCCEED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("UPDATE PASSWORD FOR USERNAME: %s BY ADMIN USER %s - %s FAILED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminUpdateUserEmailAttempt(request *http.Request, _ time.Time, usernameOrCookies, targetUsername string, succeed bool) error {
	if succeed {
		log.Printf("UPDATE EMAIL FOR USERNAME: %s BY ADMIN USER %s - %s SUCCEED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("UPDATE EMAIL FOR USERNAME: %s BY ADMIN USER %s - %s FAILED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminCreateUserAttempt(request *http.Request, _ time.Time, usernameOrCookies, targetUsername string, succeed bool) error {
	if succeed {
		log.Printf("USER CREATION: %s BY ADMIN USER %s - %s SUCCEED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("USER CREATION: %s BY ADMIN USER %s - %s FAILED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminDisableUserAttempt(request *http.Request, _ time.Time, usernameOrCookies, targetUsername string, succeed bool) error {
	if succeed {
		log.Printf("USER DISABLED: %s BY ADMIN USER %s - %s SUCCEED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("USER DISABLED: %s BY ADMIN USER %s - %s FAILED\n", targetUsername, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminAddTerminalsAttempt(request *http.Request, _ time.Time, usernameOrCookies string, terminals []string, succeed bool) error {
	if succeed {
		log.Printf("ADD TERMINALS: %s BY ADMIN USER %s - %s SUCCEED\n", terminals, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("ADD TERMINALS: %s BY ADMIN USER %s - %s FAILED\n", terminals, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogAdminAddRoutesAttempt(request *http.Request, _ time.Time, usernameOrCookies string, numberOfRoutes int, succeed bool) error {
	if succeed {
		log.Printf("ADD %d ROUTES BY ADMIN USER %s - %s SUCCEED\n", numberOfRoutes, usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("ADD %d ROUTES BY ADMIN USER %s - %s FAILED\n", numberOfRoutes, usernameOrCookies, request.RemoteAddr)
	}
	return nil
}

func (c *Callbacks) LogListTerminalsAttempt(request *http.Request, _ time.Time, usernameOrCookies string, succeed bool) error {
	if succeed {
		log.Printf("LIST TERMINALS: %s BY USER %s - %s SUCCEED\n", usernameOrCookies, request.RemoteAddr)
	} else {
		log.Printf("LIST TERMINALS: %s BY USER %s - %s FAILED\n", usernameOrCookies, request.RemoteAddr)
	}
	return nil
}
