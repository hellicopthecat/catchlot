package handler

import (
	"context"
	"fmt"
	"strconv"

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

	q := c.Queries()
	rankVal := q["rank"]
	roundIdVal := q["roundId"]
	rank, err := strconv.Atoi(rankVal)
	if err != nil {
		return commons.CheckAtoi(c, err)
	}
	round, err := strconv.Atoi(roundIdVal)
	if err != nil {
		return commons.CheckAtoi(c, err)
	}

	var dto request.CreateUserTicketRequest

	for i := 1; i <= 6; i++ {
		key := fmt.Sprintf("num_%d", i)
		num := q[key]
		convNum, err := strconv.Atoi(num)
		if err != nil {
			return commons.CheckAtoi(c, err)
		}
		dto.Number = append(dto.Number, convNum)
	}

	dto.User_id = claims.Email
	dto.Rank = rank
	dto.Round_id = round

	h.ticketService.SCreateUserTickets(ctx, dto)

	return c.Status(fiber.StatusOK).JSON(commons.Results{})
}
