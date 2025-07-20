package user

import (
	"url_shorter_new/internal/db"

	"gorm.io/gorm"
)

var (
	UsersDB *gorm.DB
	err     error
)

func InitDB() {
	UsersDB, err = db.Open("internal/db/storage/Users.db", &gorm.Config{}, &User{})
	if err != nil {
		panic(err)
	}
}
