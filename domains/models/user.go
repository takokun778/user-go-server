package models

import "github.com/google/uuid"

type User struct {
	id   string
	name string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func NewUser(name string) (user *User) {
	user = new(User)
	uuid, _ := uuid.NewRandom()
	user.id = uuid.String()
	user.name = name
	return
}
