package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title,omitempty" gorm:"unique"`
	Author string `json:"author,omitempty"`
	Rating int    `json:"rating,omitempty"`
}
