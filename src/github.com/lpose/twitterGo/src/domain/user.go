package domain

type User struct {
	id       int
	Name     string
	mail     string
	password string
	Nick     string
}

func NewUser(id int, name, mail, password, nick string) *User {
	user := User{
		id,
		name,
		mail,
		password,
		nick,
	}

	return &user
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) GetPass() string {
	return u.password
}
