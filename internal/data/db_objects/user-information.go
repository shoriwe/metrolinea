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
	Email             string
	EmergencyContact  string
	BirthDate         time.Time
}

type MetrolineaCardInformation struct {
	Id     uint
	Number string
}

type Whoami struct {
	UserId           uint
	Kind             uint
	Username         string
	Name             string
	BirthDate        time.Time
	CardNumber       string
	Email            string
	EmergencyContact string
}
