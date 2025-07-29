package url

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	UserID      uint `gorm:"index"`
	OriginalURL string
	ShortCode   string `gorm:"unique"`
}
