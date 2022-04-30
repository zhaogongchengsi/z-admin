package controller

type User struct {
}

func (u *User) Register() (user *User, err error) {
	return &User{}, nil
}
