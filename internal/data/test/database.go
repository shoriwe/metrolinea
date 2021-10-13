package test

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/errors"
	"log"
	"time"
)

var userSessions = map[string]db_objects.UserInformation{}

func LogError(_ time.Time, err error) error {
	log.Println(err)
	return nil
}

func LogLoginAttempt(_ time.Time, userInformation db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Printf("LOGIN FOR USER: %s <%d> SUCCEED\n", userInformation.Username, userInformation.Id)
	} else {
		log.Printf("LOGIN FOR USER: %s <%d> FAILED\n", userInformation.Username, userInformation.Id)
	}
	return nil
}

func LogLogoutAttempt(_ time.Time, userInformation db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Printf("LOGOUT FOR USER: %s <%d> SUCCEED\n", userInformation.Username, userInformation.Id)
	} else {
		log.Printf("LOGOUT FOR USER: %s <%d> FAILED\n", userInformation.Username, userInformation.Id)
	}
	return nil
}

func LogWhoamiAttempt(_ time.Time, whoami db_objects.Whoami, succeed bool) error {
	if succeed {
		log.Printf("WHOAMI FOR USER: %s <%d> SUCCEED\n", whoami.Username, whoami.UserId)
	} else {
		log.Printf("WHOAMI FOR USER: %s <%d> FAILED\n", whoami.Username, whoami.UserId)
	}
	return nil
}

func LogCookieGenerationAttempt(_ time.Time, userInformation db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Printf("COOKIE GENERATION FOR USER: %s <%d> SUCCEED\n", userInformation.Username, userInformation.Id)
	} else {
		log.Printf("COOKIE GENERATION FOR USER: %s <%d> FAILED\n", userInformation.Username, userInformation.Id)
	}
	return nil
}

func GenerateCookie(userInformation db_objects.UserInformation) (string, error) {
	for i := 0; i < 5; i++ {
		rawUniqueId := make([]byte, 64)
		_, readRandError := rand.Read(
			rawUniqueId)
		if readRandError != nil {
			return "", readRandError
		}
		uniqueId := base32.HexEncoding.EncodeToString(rawUniqueId)
		_, ok := userSessions[uniqueId]
		if !ok {
			userSessions[uniqueId] = userInformation
			return uniqueId, nil
		}
	}
	return "", errors.CookieGenerationError(userInformation)
}

func Login(username, password string) (db_objects.UserInformation, bool, error) {
	if username == "John" && password == "Hasta la vista baby!" {
		return db_objects.UserInformation{
			Id:           1,
			Kind:         db_objects.Administrator,
			Username:     "terminator",
			PasswordHash: "",
		}, true, nil
	} else if username == "Marla" && password == "Tyler" {
		return db_objects.UserInformation{
			Id:           2,
			Kind:         db_objects.User,
			Username:     "mSinger",
			PasswordHash: "",
		}, true, nil
	}
	return db_objects.UserInformation{}, false, nil
}

func Logout(cookies string) (db_objects.UserInformation, bool, error) {
	userInformation, found := userSessions[cookies]
	if !found {
		return db_objects.UserInformation{}, false, nil
	}
	delete(userSessions, cookies)
	return userInformation, true, nil
}

func Whoami(cookies string) (db_objects.Whoami, bool, error) {
	userInformation, found := userSessions[cookies]
	if !found {
		return db_objects.Whoami{}, false, nil
	}
	if userInformation.Id == 1 { // John Connor
		return db_objects.Whoami{
			UserId:     userInformation.Id,
			Kind:       userInformation.Kind,
			Username:   userInformation.Username,
			Name:       "John Connor",
			BirthDate:  time.Now(),
			CardNumber: "34252350",
		}, true, nil
	} else if userInformation.Id == 2 { // Marla Singer
		return db_objects.Whoami{
			UserId:     userInformation.Id,
			Kind:       userInformation.Kind,
			Username:   userInformation.Username,
			Name:       "Marla Singer",
			BirthDate:  time.Now(),
			CardNumber: "92378457",
		}, true, nil
	}
	return db_objects.Whoami{}, false, nil
}
