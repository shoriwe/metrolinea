package test

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/errors"
)

var userSessions = map[string]*db_objects.UserInformation{}

func CheckCookies(cookies string) (*db_objects.UserInformation, bool, error) {
	userInformation, found := userSessions[cookies]
	return userInformation, found, nil
}

func GenerateCookie(userInformation *db_objects.UserInformation) (string, error) {
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

func Logout(cookies string) (bool, error) {
	_, found := userSessions[cookies]
	if !found {
		return false, nil
	}
	delete(userSessions, cookies)
	return true, nil
}
