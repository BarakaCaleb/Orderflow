package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"fullname"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
