package repo

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (r TicketRepo) RFindUserTickets(c context.Context, req request.TFindUserTicketsRequest) ([]request.TFindUserTicketResponse, error) {
	tx, err := r.db.BeginTx(c, nil)
	if err != nil {
		log.Printf("RFindUserTickets TX 시작 실패 :: %s", err)
		return nil, err
	}
	defer tx.Rollback()

	q, err := os.ReadFile(constants.SelectSQL + "user_tickets" + "/s_findUserTickets.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}

	q2, err := os.ReadFile(constants.SelectSQL + "user_tickets" + "/s_findUserTicketNumberById.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}

	rows, err := tx.QueryContext(c, string(q), req.Id, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}

	var tickets []request.TFindUserTicketResponse
	for rows.Next() {

		var result request.TFindUserTicketResponse
		err := rows.Scan(
			&result.Id,
			&result.Created_at,
			&result.Updated_at,
			&result.User_id,
			&result.Round_id,
			&result.Rank,
			&result.Bonus_match,
			&result.Checked,
		)
		if err != nil {
			rows.Close()
			return nil, err
		}
		tickets = append(tickets, result)
	}
	if len(tickets) == 0 {
		return []request.TFindUserTicketResponse{}, errors.New("조회결과가 없습니다.")
	}
	for i := range tickets {
		numRows, err := tx.QueryContext(c, string(q2), tickets[i].Id)
		if err != nil {
			return nil, err
		}
		var pickNum []request.TPickNumberResponse
		for numRows.Next() {
			var pn request.TPickNumberResponse
			if err := numRows.Scan(
				&pn.Pick_number,
				&pn.Pick_number_id,
			); err != nil {
				numRows.Close()
				return nil, err
			}
			pickNum = append(pickNum, pn)
		}
		numRows.Close()
		tickets[i].Numbers = pickNum
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return tickets, nil
}
