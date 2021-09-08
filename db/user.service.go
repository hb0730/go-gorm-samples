package db

import "fmt"

type UserService struct {
	repository *UserRepository
}

func NewService() UserService {
	return UserService{
		repository: NewUserRepository(),
	}
}

func (service UserService) Insert(user *User) {
	_, err := service.repository.Insert(&user, DB)
	if err != nil {
		fmt.Printf("insert user failed:%v", err)
	}
}
