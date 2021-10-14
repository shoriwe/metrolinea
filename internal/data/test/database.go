package test

import (
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"time"
)

var (
	lastUserId    uint = 2
	usersDatabase      = map[string]*db_objects.UserInformation{
		"terminator": {
			Id:               1,
			Kind:             db_objects.Administrator,
			Username:         "terminator",
			PasswordHash:     "Hasta la vista baby!",
			Name:             "John Connor",
			Email:            "jonny@skynet.corp",
			EmergencyContact: "MOTHER_CELLPHONE_HERE",
			BirthDate:        time.Time{},
		},
		"mSinger": {
			Id:               2,
			Kind:             db_objects.User,
			Username:         "mSinger",
			PasswordHash:     "The first rule of the fight club is...",
			Name:             "Marla Singer",
			Email:            "marla@paper.street",
			EmergencyContact: "TYLER_PHONE_HERE",
			BirthDate:        time.Time{},
		},
	}
)

func Register(registrationForm *forms.RegisterForm) (bool, string, error) {
	// ToDo: Do something to sanitize input
	lastUserId++
	usersDatabase[registrationForm.Username] = &db_objects.UserInformation{
		Id:               lastUserId,
		Kind:             db_objects.User,
		Username:         registrationForm.Username,
		PasswordHash:     registrationForm.Password,
		Name:             registrationForm.Name,
		Email:            registrationForm.Email,
		EmergencyContact: registrationForm.EmergencyContact,
		BirthDate:        registrationForm.BirthDate,
	}
	return true, "Successfully created user", nil
}

func Login(username, password string) (*db_objects.UserInformation, bool, error) {
	userInformation, userFound := usersDatabase[username]
	if !userFound {
		return nil, false, nil
	}
	if userInformation.PasswordHash == password {
		return userInformation, true, nil
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
