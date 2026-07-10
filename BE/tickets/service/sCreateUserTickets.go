package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (s TicketService) SCreateUserTickets(ctx context.Context, req request.CreateUserTicketRequest) error {
	return s.repo.RCreateUserTickets(ctx, req)
}
