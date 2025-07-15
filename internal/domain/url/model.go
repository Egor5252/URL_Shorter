package url

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	UserID      uint
	OriginalURL string
	ShortCode   string `gorm:"unique"`
	Count       uint32
}
