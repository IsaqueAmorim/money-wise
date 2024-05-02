package user

import (
	"time"

	passwordHelper "github.com/IsaqueAmorim/money-wise/src/domain/user/password-helper"
)

type User struct {
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func NewUser(username, email, password string) *User {
	return &User{
		Username:  username,
		Email:     email,
		Password:  HashPassword(password),
		CreatedAt: time.Now(),
	}
}

func (u User) CheckPasswordIsValid(password string) bool {
	return passwordHelper.CheckPasswordHash(password, u.Password)
}

func HashPassword(pass string) string {
	password, _ := passwordHelper.HashPassword(pass)
	return password
}
