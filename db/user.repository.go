package db

type UserRepository struct {
	BaseMapper
}

func NewUserRepository() *UserRepository {
	repository := new(UserRepository)
	return repository
}
