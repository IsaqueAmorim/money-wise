package entity

import "time"

type User struct {
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}
