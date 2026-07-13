package service

import (
	"context"

	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (s TicketService) SCreateUserTickets(ctx context.Context, id string, dto []request.CreateUserTicketDto) error {
	var req request.CreateUserTicketRequest
	req.User_id = id
	req.Ticket_info = dto
	return s.repo.RCreateUserTickets(ctx, req)
}
