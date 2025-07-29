package login

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	UserID           string `gorm:"not null"`
	RefreshTotenHash string `gorm:"not null"`
}
