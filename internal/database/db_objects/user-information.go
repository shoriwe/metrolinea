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
}

type PersonInformation struct {
	Id                uint
	UserInformationId uint
	MetrolineaCardId  uint
	Name              string
	BirthDate         time.Time
}

type MetrolineaCardInformation struct {
	Id     uint
	Number string
}
