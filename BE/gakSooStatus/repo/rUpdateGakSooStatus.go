package repo

import (
	"context"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/gakSooStatus/request"
)

func (r GakSooStatusRepo) RUpdateGakSooStatus(ctx context.Context, dto request.GakSooStatusUpdateRequest) error {
	q, err := os.ReadFile(constants.UpdateSQL + "u_gak_soo_status.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return err
	}

	result, err := r.db.ExecContext(ctx, string(q), dto.Bonus_count, dto.Soo_id)
	if err != nil {
		log.Printf("❌ Update GakSooStatus Counts :: %s", err)
		return err
	}

	if _, err := result.RowsAffected(); err != nil {
		log.Printf("❌ Update GakSooStatus Counts :: %s", err)
		return err
	}

	return nil
}
