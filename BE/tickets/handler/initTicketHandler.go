package handler

import (
	"github.com/hellicopthecat/catchlot/tickets/service"
)

type TicketHandler struct {
	ticketService *service.TicketService
}

func InitTicketHandler(service *service.TicketService) *TicketHandler {
	return &TicketHandler{
		ticketService: service,
	}
}
