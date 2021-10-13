package test

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/shoriwe/metrolinea/internal/database/db_objects"
	"github.com/shoriwe/metrolinea/internal/errors"
	"log"
	"time"
)

var cookies = map[string]db_objects.UserInformation{}

func LogError(_ time.Time, err error) error {
	log.Println(err)
	return nil
}

func LogLoginAttempt(_ time.Time, userInformation db_objects.UserInformation, succeed bool) error {
	if succeed {
		log.Println("LOGIN SUCCEED FOR:", userInformation)
	} else {
		log.Println("LOGIN FAILED FOR:", userInformation)
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
		_, ok := cookies[uniqueId]
		if !ok {
			cookies[uniqueId] = userInformation
			return uniqueId, nil
		}
	}
	return "", errors.CookieGenerationError(userInformation)
}

func Login(username, password string) (db_objects.UserInformation, bool, error) {
	if username == "John" && password == "Connor" {
		return db_objects.UserInformation{
			Id:           1,
			Kind:         db_objects.Administrator,
			Username:     "terminator",
			PasswordHash: "",
		}, true, nil
	} else if username == "Marla" && password == "Singer" {
		return db_objects.UserInformation{
			Id:           2,
			Kind:         db_objects.User,
			Username:     "mSinger",
			PasswordHash: "",
		}, true, nil
	}
	return db_objects.UserInformation{}, false, nil
}
