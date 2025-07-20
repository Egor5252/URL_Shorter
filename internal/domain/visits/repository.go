package visits

import (
	"url_shorter_new/internal/db"

	"gorm.io/gorm"
)

var (
	VisitsDB *gorm.DB
	err      error
)

func InitDB() {
	VisitsDB, err = db.Open("internal/db/storage/VisitsDB.db", &gorm.Config{}, &Visits{})
	if err != nil {
		panic(err)
	}
}
