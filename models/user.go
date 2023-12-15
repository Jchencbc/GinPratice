package models

type Auth_user struct {
	Id           int
	Password     string
	is_superuser int
	username     string
}
