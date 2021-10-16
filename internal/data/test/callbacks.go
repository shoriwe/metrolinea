package test

import (
	"github.com/shoriwe/metrolinea/internal/data/db_objects"
	"time"
)

type Callbacks struct {
	userSessions  map[string]*db_objects.UserInformation
	lastUserId    uint
	usersDatabase map[string]*db_objects.UserInformation
}

func NewCallbacks() *Callbacks {
	return &Callbacks{
		userSessions: map[string]*db_objects.UserInformation{},
		lastUserId:   2,
		usersDatabase: map[string]*db_objects.UserInformation{
			"terminator": {
				Id:           1,
				Kind:         db_objects.Administrator,
				Username:     "terminator",
				PasswordHash: "Hasta la vista baby!",
				Name:         "John Connor",
				Email:        "jonny@skynet.corp",
				BirthDate:    time.Time{},
			},
			"mSinger": {
				Id:           2,
				Kind:         db_objects.User,
				Username:     "mSinger",
				PasswordHash: "The first rule of the fight club is...",
				Name:         "Marla Singer",
				Email:        "marla@paper.street",
				BirthDate:    time.Time{},
			},
		},
	}
}
