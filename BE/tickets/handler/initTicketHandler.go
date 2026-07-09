package handler

import (
	"github.com/gofiber/fiber/v3"
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

func (t TicketHandler) HCreateUserTicket(c fiber.Ctx) error {
	return nil
}
