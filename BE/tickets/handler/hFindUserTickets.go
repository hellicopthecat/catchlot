package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (h TicketHandler) HFindUserTickets(c fiber.Ctx) error {
	ctx := context.Background()
	at := c.Cookies(constants.ACCESS)

	limit := c.Query("limit", "10")
	offset := c.Query("offset", "0")

	l, err := strconv.Atoi(limit)
	if err != nil {
		return commons.CheckAtoi(c, err)
	}
	o, err := strconv.Atoi(offset)
	if err != nil {
		return commons.CheckAtoi(c, err)
	}

	claims, err := commons.DecodedACCESSJWT(at)
	if err != nil {
		return commons.UnauthorizedError(c)
	}
	var req request.TFindUserTicketsRequest
	req.Id = claims.Email
	req.Limit = l
	req.Offset = o

	tickets, err := h.ticketService.SFindUserTickets(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(commons.Results{
			Status: false,
			Msg:    "유저의 구매이력을 조회하지 못했습니다.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(commons.Results{
		Status: true,
		Msg:    "Success",
		Data:   tickets,
	})
}
