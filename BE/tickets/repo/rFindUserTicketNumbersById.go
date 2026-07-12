package repo

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (r TicketRepo) RFindUserTicketById(ctx context.Context, req request.TFindUserTicketByIdRequest) (*request.TFindUserTicketResponse, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("RFindUserTicketById TX 시작 실패 :: %s", err)
		return nil, err
	}
	defer tx.Rollback()
	q, err := os.ReadFile(constants.SelectSQL + "/s_findUserTicket.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}

	var result request.TFindUserTicketResponse
	err = tx.QueryRow(string(q), req.Id).Scan(
		&result.Id,
		&result.Created_at,
		&result.Updated_at,
		&result.User_id,
		&result.Round_id,
		&result.Rank,
		&result.Bonus_match,
		&result.Checked)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("조회결과가 없습니다.")
		}
		return nil, err
	}

	q2, err := os.ReadFile(constants.SelectSQL + "/s_findUserTicketNumberById.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}

	rows, err := tx.QueryContext(ctx, string(q2), result.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("조회결과가 없습니다.")
		}
		return nil, err
	}

	for rows.Next() {
		var ticketNum request.TPickNumberResponse
		if err := rows.Scan(
			&ticketNum.Pick_number,
			&ticketNum.Pick_number_id,
		); err != nil {
			rows.Close()
			return nil, err
		}
		result.Numbers = append(result.Numbers, ticketNum)
	}
	tx.Commit()
	return &result, nil
}
