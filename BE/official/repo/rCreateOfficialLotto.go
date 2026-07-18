package repo

import (
	"context"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/official/request"
)

func (r OfficialLottoRepo) RCreateOfficialLotto(ctx context.Context, dto request.LottoRoundRequest) error {
	q1, err := os.ReadFile(constants.InsertSQL + "i_lotto_rounds.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return err
	}

	q2, err := os.ReadFile(constants.InsertSQL + "i_lotto_rounds_numbers.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return err
	}
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("❌ RCreateOfficialLotto :: %s ", err)
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, string(q1), dto.Id, dto.Round_no, dto.Draw_date, dto.Bonus_number)
	if err != nil {
		return err
	}

	for _, nums := range dto.Numbers {
		_, err = tx.ExecContext(ctx, string(q2), dto.Id, nums)
		if err != nil {
			return err
		}
	}
	tx.Commit()

	return nil
}
