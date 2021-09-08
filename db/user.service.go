package db

import "fmt"

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
