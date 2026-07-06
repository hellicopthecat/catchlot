package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/users/request"
)

func (s UserService) SUpdateUser(ctx context.Context, req request.UpdateUserRequest) error {
	return s.repo.RUpdateUser(ctx, req)
}
