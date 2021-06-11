package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string `json:"uname"`
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
