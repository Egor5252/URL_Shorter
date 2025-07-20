package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	PassHash string `gorm:"not null"`
}
