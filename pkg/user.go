package pkg

type User struct {
	ID         string `json:ID`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Picture    string `json:"picture"`
	IsLoggedIn bool   `json:"loggedIn"`
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
	if usr.ID == "" {
		panic("User Id is not found")
	}
	r.Users[usr.ID] = usr
}

func (r *UserConstractor) GetUsers() map[string]User {
	return r.Users
}

func (r *UserConstractor) DeleteUser(id string) {
	delete(r.Users, id)
}
