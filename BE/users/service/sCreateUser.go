package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/users/request"
)

func (s UserService) SCreateUser(ctx context.Context, userReq request.CreateUserRequest) error {
	return s.repo.RCreateUser(ctx, userReq)
}
