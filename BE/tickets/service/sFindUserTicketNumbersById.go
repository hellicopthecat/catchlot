package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (s TicketService) SFindUserTicketNumbers(ctx context.Context, req request.TFindUserTicketByIdRequest) (*request.TFindUserTicketResponse, error) {
	return s.repo.RFindUserTicketById(ctx, req)
}
