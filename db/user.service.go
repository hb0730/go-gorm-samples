package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	repository *UserRepository
}

func NewService() UserService {
	return UserService{
		repository: NewUserRepository(DB),
	}
}

func (service UserService) Insert(user *User) {
	_, err := service.repository.Insert(&user)
	if err != nil {
		fmt.Printf("insert user failed:%v", err)
	}
}

func (service UserService) TransactionTest() {
	err := DB.Transaction(func(tx *gorm.DB) error {
		repository := NewUserRepository(tx)
		user := User{Name: "事务回滚", Age: 110}
		repository.Insert(&user)
		return errors.New("测试回滚")
	})
	if err != nil {
		fmt.Printf("insert user failed:%v", err)
	}
}
