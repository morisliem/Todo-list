package src

import "time"

type User struct {
	Id         string `json:Id`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
	Created_at time.Time
	Deleted_at time.Time
}

type UserConstractor struct {
	Users map[string]User
}

func NewUser() *UserConstractor {
	return &UserConstractor{
		Users: make(map[string]User),
	}
}

func (r *UserConstractor) AddUser(usr User) {
	if usr.Id == "" {
		panic("User Id is not found")
	}
	r.Users[usr.Id] = usr
}

func (r *UserConstractor) GetUsers() map[string]User {
	return r.Users
}

func (r *UserConstractor) DeleteUser(id string) {
	delete(r.Users, id)
}
