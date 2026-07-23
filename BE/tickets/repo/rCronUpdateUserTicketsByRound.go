package repo

import (
	"context"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/tickets/request"
)

func (r TicketRepo) RCronUpdateUserTicketsByRound(ctx context.Context, dto request.UpdateUserTicketRoundRequest) error {
	q, err := os.ReadFile(constants.UpdateSQL + "u_users_ticketsByRound.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("❌ RCronUpdateUserTicketsByRound tx 실패 :: %s ", err)
		return err
	}
	defer tx.Rollback()
	// 해당 회차 의 유저 티켓을 전부 가져오고 유저아이디를 가져온다. 배열이 0개면 패스
	// 해당 회차의 유저 티켓에 등급 / 보너스 매치 / 크론 체크 값에 값을 넣어준다.
	// 해당 회차의 유저 티켓의 유저 아이디를 가져와 for문을 돌려 user_rate에 user_id와 매치시켜 값을 넣어준다.
	tx.Commit()
	return nil
}
