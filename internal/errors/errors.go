package errors

import (
	"errors"
	"fmt"
	"metrolinea/internal/database/db_objects"
)

func CookieGenerationError(userInformation db_objects.UserInformation) error {
	return errors.New(fmt.Sprintf("Cookie generation error for: %v", userInformation))
}

func GoRuntimeError(err error, remoteAddress, method, uri string) error {
	return errors.New(fmt.Sprintf("%s %s %s -> %s", method, uri, remoteAddress, err))
}

func MethodNotAllowed(remoteAddress, method, uri string) error {
	return errors.New(fmt.Sprintf("%s %s FROM %s NOT ALLOWED", method, uri, remoteAddress))
}

func ContentTypeNotSupported(remoteAddress, method, uri, contentType string) error {
	return errors.New(fmt.Sprintf("%s %s %s FROM %s NOT SUPPORTED", method, uri, contentType, remoteAddress))
}
