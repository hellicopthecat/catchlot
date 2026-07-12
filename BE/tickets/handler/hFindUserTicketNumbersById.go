package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (h TicketHandler) HFindUserTicketNumber(c fiber.Ctx) error {
	ctx := context.Background()
	tId := c.Query("ticket_id")
	id, err := strconv.Atoi(tId)
	if err != nil {
		log.Printf("❌ HFindUserTicketNumber :: %s", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(commons.Results{
			Status: false,
			Msg:    "잘못된 요청입니다.",
		})
	}
	var req request.TFindUserTicketByIdRequest
	req.Id = id

	ticket, err := h.ticketService.SFindUserTicketNumbers(ctx, req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(commons.Results{
			Status: false,
			Msg:    "유저의 구매이력을 조회하지 못했습니다.",
		})
	}
	return c.Status(fiber.StatusOK).JSON(commons.Results{
		Status: true,
		Msg:    "Success",
		Data:   ticket,
	})
}
