package repo

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
)

func (repo OfficialLottoRepo) RFindOfficialLottoByRound(ctx context.Context, id int) (*string, error) {
	q, err := os.ReadFile(constants.SelectSQL + "lotto" + "/s_findLottoRoundsByRound.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return nil, err
	}
	var lotto_id string
	err = repo.db.QueryRow(string(q), id).Scan(&lotto_id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Printf("❌ RFindOfficialLottoByRound :: %s", err)
		return nil, err
	}

	return &lotto_id, nil
}
