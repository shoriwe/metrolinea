package test

import (
	"github.com/shoriwe/metrolinea/internal/api/forms"
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"net/http"
)

func (c *Callbacks) Register(_ *http.Request, registrationForm *forms.RegisterForm) (bool, string, error) {
	// ToDo: Do something to sanitize input
	c.lastUserId++
	c.usersDatabase[registrationForm.Username] = &db_objects.UserInformation{
		Id:           c.lastUserId,
		Kind:         db_objects.User,
		Username:     registrationForm.Username,
		PasswordHash: registrationForm.Password,
		Name:         registrationForm.Name,
		Email:        registrationForm.Email,
		BirthDate:    registrationForm.BirthDate,
	}
	return true, "Successfully created user", nil
}

func (c *Callbacks) Login(request *http.Request, username, password string) (*db_objects.UserInformation, bool, error) {
	userInformation, userFound := c.usersDatabase[username]
	if !userFound {
		return nil, false, nil
	}
	if userInformation.PasswordHash == password {
		return userInformation, true, nil
	}
	return nil, false, nil
}

func (c *Callbacks) CheckUserExists(_ *http.Request, username string) (bool, error) {
	switch username {
	case "mSinger", "terminator":
		return true, nil
	}
	return false, nil
}

func (c *Callbacks) UpdatePassword(_ *http.Request, username string, oldPassword, newPassword string) (bool, string, error) {
	userInformation, found := c.usersDatabase[username]
	if !found {
		return false, "Invalid credentials", nil
	}
	if userInformation.PasswordHash != oldPassword {
		return false, "Invalid credentials", nil
	}
	if oldPassword == newPassword {
		return false, "New password is equal to old", nil
	}
	// ToDo: Check if new password is weak
	c.usersDatabase[username].PasswordHash = newPassword
	return true, "Update Succeed", nil
}

func (c *Callbacks) UpdateEmail(_ *http.Request, username string, password, newEmail string) (bool, string, error) {
	userInformation, found := c.usersDatabase[username]
	if !found {
		return false, "Invalid credentials", nil
	}
	if userInformation.PasswordHash != password {
		return false, "Invalid credentials", nil
	}
	if userInformation.Email == newEmail {
		return false, "New email is equal to old", nil
	}
	// ToDo: Check if new email is valid
	c.usersDatabase[username].Email = newEmail
	return true, "Update Succeed", nil
}
