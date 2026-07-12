package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (s TicketService) SFindUserTickets(ctx context.Context, req request.TFindUserTicketsRequest) ([]request.TFindUserTicketResponse, error) {
	return s.repo.RFindUserTickets(ctx, req)
}
