package repo

import (
	"context"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
)

func (repo OfficialLottoRepo) RFindOfficialLottoByRound(ctx context.Context, id int) (*int, error) {
	q, err := os.ReadFile(constants.SelectSQL + "s_findLottoRoundsByRound")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}
	row := repo.db.QueryRow(string(q), id)
	if row.Err() != nil {
		log.Printf("❌ RFindOfficialLottoByRound :: %s", err)
		return nil, err
	}
	var lotto_id int
	row.Scan(&lotto_id)

	return &lotto_id, nil
}
