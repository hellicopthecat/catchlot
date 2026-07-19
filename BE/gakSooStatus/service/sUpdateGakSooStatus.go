package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/gakSooStatus/request"
)

func (s GakSooStatusService) SUpdateGakSooStatus(ctx context.Context, dto request.GakSooStatusUpdateRequest) error {
	return s.repo.RUpdateGakSooStatus(ctx, dto)
}
