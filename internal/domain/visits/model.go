package visits

import "gorm.io/gorm"

type Visits struct {
	gorm.Model
	ShortUrlID uint `gorm:"index"`
	UserIP     string
}
