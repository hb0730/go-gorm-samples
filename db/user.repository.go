package db

import "gorm.io/gorm"

type UserRepository struct {
	BaseMapper
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	repository := new(UserRepository)
	repository.DB = db
	return repository
}
