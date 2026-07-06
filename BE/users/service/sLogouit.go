package service

import (
	"context"
)

func (s *UserService) SLogout(ctx context.Context, email string) error {
	return s.repo.RLogout(ctx, email)
}
