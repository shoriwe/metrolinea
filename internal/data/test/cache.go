package test

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"github.com/shoriwe/metrolinea/internal/errors"
	"net/http"
)

func (c *Callbacks) CheckCookies(_ *http.Request, cookies string) (*db_objects.UserInformation, bool, error) {
	userInformation, found := c.userSessions[cookies]
	return userInformation, found, nil
}

func (c *Callbacks) GenerateCookie(_ *http.Request, userInformation *db_objects.UserInformation) (string, error) {
	for i := 0; i < 5; i++ {
		rawUniqueId := make([]byte, 64)
		_, readRandError := rand.Read(
			rawUniqueId)
		if readRandError != nil {
			return "", readRandError
		}
		uniqueId := base32.HexEncoding.EncodeToString(rawUniqueId)
		_, ok := c.userSessions[uniqueId]
		if !ok {
			c.userSessions[uniqueId] = userInformation
			return uniqueId, nil
		}
	}
	return "", errors.CookieGenerationError(userInformation)
}

func (c *Callbacks) Logout(_ *http.Request, cookies string) (bool, error) {
	_, found := c.userSessions[cookies]
	if !found {
		return false, nil
	}
	delete(c.userSessions, cookies)
	return true, nil
}
