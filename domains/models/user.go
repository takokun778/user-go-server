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

func newUser(id string, name string) (user *User) {
	user = new(User)
	user.id = id
	user.name = name
	return
}

func CreateNewUser(name string) (user *User) {
	uuid, _ := uuid.NewRandom()
	user = newUser(uuid.String(), name)
	return
}

func CreatePlainUser(id string, name string) (user *User) {
	user = newUser(id, name)
	return
}
