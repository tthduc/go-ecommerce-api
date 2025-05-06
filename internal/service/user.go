package service

import "go-ecommerce-api/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) Register() string {
	return us.UserRepo.Register()
}
