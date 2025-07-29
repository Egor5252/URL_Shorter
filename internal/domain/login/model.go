package login

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	User_ID          string `gorm:"not null"`
	RefreshTokenHash string `gorm:"not null"`
	User_IP          string `gorm:"not null"`
}
