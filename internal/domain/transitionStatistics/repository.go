package transitionstatistics

import (
	"urlShorter/internal/db"

	"gorm.io/gorm"
)

var (
	TSDB *gorm.DB
	err  error
)

func InitDB() {
	TSDB, err = db.Open("internal/db/storage/TSDB.db", &gorm.Config{}, &Transitionstatistics{})
	if err != nil {
		panic(err)
	}
}
