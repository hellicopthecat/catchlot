package service

import "github.com/hellicopthecat/catchlot/users/repo"

type UserService struct {
	repo repo.UserRepo
}

func InitUserService(repo repo.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}
