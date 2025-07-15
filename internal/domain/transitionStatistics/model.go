package transitionstatistics

import "gorm.io/gorm"

type Transitionstatistics struct {
	gorm.Model
	UserIP   string
	ShortUrl string
}
