package url

import (
	"urlShorter/internal/db"

	"gorm.io/gorm"
)

var (
	UrlDB *gorm.DB
	err   error
)

func InitDB() {
	UrlDB, err = db.Open("internal/db/storage/UrlDB.db", &gorm.Config{}, &Url{})
	if err != nil {
		panic(err)
	}
}
