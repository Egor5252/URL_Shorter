package db

import (
	"fmt"

	"gorm.io/gorm"
)

// CRUD
func Create[T any](db *gorm.DB, entity *T) error {
	if db == nil {
		return fmt.Errorf("БД не инициализирована")
	}

	if err := db.Create(entity).Error; err != nil {
		return err
	}

	return nil
}

func ReadByID[T any](db *gorm.DB, id uint) (*T, error) {
	if db == nil {
		return nil, fmt.Errorf("БД не инициализирована")
	}

	var entity T
	if err := db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func ReadFirstByValue[T any](db *gorm.DB, field string, value any) (*T, error) {
	if db == nil {
		return nil, fmt.Errorf("БД не инициализирована")
	}

	var entity T
	if err := db.First(&entity, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func ReadOneByValues[T any](db *gorm.DB, conditions map[string]any) (*T, error) {
	if db == nil {
		return nil, fmt.Errorf("БД не инициализирована")
	}

	var entity T
	query := db
	for field, value := range conditions {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	if err := query.First(&entity).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func ReadAllByValue[T any](db *gorm.DB, field string, value any) ([]T, error) {
	if db == nil {
		return nil, fmt.Errorf("БД не инициализирована")
	}

	var entities []T
	if err := db.Where(fmt.Sprintf("%s = ?", field), value).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func ReadAll[T any](db *gorm.DB) ([]T, error) {
	if db == nil {
		return nil, fmt.Errorf("БД не инициализирована")
	}

	var entities []T
	if err := db.Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func Update[T any](db *gorm.DB, entity *T, field string, val any) error {
	if db == nil {
		return fmt.Errorf("БД не инициализирована")
	}

	return db.Model(entity).Update(field, val).Error
}

func Delete[T any](db *gorm.DB, entity *T) error {
	if db == nil {
		return fmt.Errorf("БД не инициализирована")
	}

	return db.Delete(entity).Error
}
