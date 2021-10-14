package test

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"time"
)

func Login(username, password string) (*db_objects.UserInformation, bool, error) {
	if username == "John" && password == "Hasta la vista baby!" {
		return &db_objects.UserInformation{
			Id:           1,
			Kind:         db_objects.Administrator,
			Username:     "terminator",
			PasswordHash: "",
		}, true, nil
	} else if username == "Marla" && password == "Tyler" {
		return &db_objects.UserInformation{
			Id:           2,
			Kind:         db_objects.User,
			Username:     "mSinger",
			PasswordHash: "",
		}, true, nil
	}
	return nil, false, nil
}

func Whoami(userInformation *db_objects.UserInformation) (*db_objects.Whoami, bool, error) {
	if userInformation.Id == 1 { // John Connor
		return &db_objects.Whoami{
			UserId:           userInformation.Id,
			Kind:             userInformation.Kind,
			Username:         userInformation.Username,
			Name:             "John Connor",
			BirthDate:        time.Now(),
			CardNumber:       "34252350",
			Email:            "johnny@skynet.corp",
			EmergencyContact: "55555555",
		}, true, nil
	} else if userInformation.Id == 2 { // Marla Singer
		return &db_objects.Whoami{
			UserId:           userInformation.Id,
			Kind:             userInformation.Kind,
			Username:         userInformation.Username,
			Name:             "Marla Singer",
			BirthDate:        time.Now(),
			CardNumber:       "92378457",
			Email:            "marla@paperstree.corp",
			EmergencyContact: "55555555",
		}, true, nil
	}
	return nil, false, nil
}

func CheckUserExists(username string) (bool, error) {
	switch username {
	case "mSinger", "terminator":
		return true, nil
	}
	return false, nil
}
