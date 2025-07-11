package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Open[T any](db_name string, config *gorm.Config, model *T) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(db_name), config)
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(model); err != nil {
		return nil, err
	}

	return db, nil
}
