package models

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}
