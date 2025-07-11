package url

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalURL string
	ShortCode   string `gorm:"unique"`
	Count       uint32
}
