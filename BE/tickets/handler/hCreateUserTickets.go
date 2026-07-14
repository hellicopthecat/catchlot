package handler

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (h TicketHandler) HCreateUserTickets(c fiber.Ctx) error {
	// TODO: body요청 및 배열 등록..
	ctx := context.Background()
	at := c.Cookies(constants.ACCESS)

	claims, err := commons.DecodedACCESSJWT(at)
	if err != nil {
		return commons.UnauthorizedError(c)
	}

	var dto []request.CreateUserTicketDto

	if err := c.Bind().Body(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(commons.Results{
			Status: false,
			Msg:    "올바른 요청이 아닙니다.",
		})
	}

	log.Printf("tickets :: %+v", dto)

	h.ticketService.SCreateUserTickets(ctx, claims.Email, dto)

	return c.Status(fiber.StatusOK).JSON(commons.Results{
		Status: true,
		Msg:    "등록성공했습니다.",
	})
}
