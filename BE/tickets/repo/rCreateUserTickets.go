package repo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (r *TicketRepo) RCreateUserTickets(ctx context.Context, req request.CreateUserTicketRequest) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("RCreateUserTicket 트렌젝션 시작 실패했습니다. :: %s", err)
		return err
	}
	defer tx.Rollback()

	q, err := os.ReadFile(constants.InsertSQL + "i_users_tickets.sql")
	if err != nil {
		log.Printf("RCreateUserTicket 파일을 읽는데 실패했습니다. :: %s", err)
		return err
	}

	q2, err := os.ReadFile(constants.InsertSQL + "i_users_tickets_numbers.sql")
	if err != nil {
		log.Printf("RCreateUserTicket ticket_num 파일을 읽는데 실패했습니다. :: %s", err)
		return err
	}

	for _, tickets := range req.Ticket_info {
		var id int
		err = tx.QueryRow(string(q), req.User_id, tickets.Round_id, tickets.Rank).Scan(&id)
		if err != nil {
			log.Printf("RCreateUserTicket 티켓을 생성하는데 실패했습니다. :: %s", err)
			return err
		}

		for _, num := range tickets.Numbers {

			gakSooId, ok := r.gak.GakSooMap[num]

			if !ok {
				return fmt.Errorf("유효하지 않은 번호 : %d", num)
			}

			_, err := tx.Exec(string(q2), id, gakSooId, num)
			if err != nil {
				log.Printf("RCreateUserTicket 티켓의 번호를 생성하는데 실패했습니다. :: %s", err)
				return err
			}
		}
	}
	return nil
}
