package db_objects

import "time"

const (
	User = iota
	Administrator
)

type UserInformation struct {
	Id           uint
	Kind         uint
	Username     string
	PasswordHash string
	Name         string
	Email        string
	BirthDate    time.Time
}
