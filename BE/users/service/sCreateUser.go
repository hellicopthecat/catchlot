package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/users/repo"
	"github.com/hellicopthecat/catchlot/users/request"
)

func SCreateUser(ctx context.Context, userRepo repo.UserRepo, userReq request.CreateUserRequest) error {
	return userRepo.RCreateUser(ctx, userReq)
}
