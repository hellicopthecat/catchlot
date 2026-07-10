package handler

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (h TicketHandler) HCreateUserTickets(c fiber.Ctx) error {
	ctx := context.Background()
	at := c.Cookies(constants.ACCESS)

	claims, err := commons.DecodedACCESSJWT(at)
	if err != nil {
		commons.UnauthorizedError(c)
	}
	q := c.Queries()
	rankVal := q["rank"]
	roundIdVal := q["roundId"]
	num_1 := q["num_1"]
	num_2 := q["num_2"]
	num_3 := q["num_3"]
	num_4 := q["num_4"]
	num_5 := q["num_5"]
	num_6 := q["num_6"]
	rank, err := strconv.Atoi(rankVal)
	round, err := strconv.Atoi(roundIdVal)

	var dto request.CreateUserTicketRequest

	dto.User_id = claims.Email
	dto.Rank = rank
	dto.Round_id = round
	append(dto.Number)

	h.ticketService.SCreateUserTickets(ctx)

	return c.Status(fiber.StatusOK).JSON(commons.Results{})
}
