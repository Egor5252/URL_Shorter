package login

import (
	"url_shorter_new/internal/db"

	"gorm.io/gorm"
)

var (
	LoginDB *gorm.DB
	err     error
)

func InitDB() {
	LoginDB, err = db.Open("internal/db/storage/LoginDB.db", &gorm.Config{}, &Login{})
	if err != nil {
		panic(err)
	}
}
