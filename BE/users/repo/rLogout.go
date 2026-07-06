package repo

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
)

func (db *UserRepo) RLogout(ctx context.Context, email string) error {
	q, err := os.ReadFile(constants.UpdateSQL + "u_users_logout.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return err
	}
	result, err := db.db.Exec(string(q), email)
	if err != nil {
		log.Printf("❌ 유저 업데이트 중 에러가 발생했습니다. %d", err)
		return err
	}
	affected, err := result.RowsAffected()
	if affected == 0 {
		log.Println("해당 유저가 존재하지 않음.")
		return fmt.Errorf("해당 유저가 존재하지 않음. :: %s ", email)
	}
	if err != nil {
		log.Printf("❌ 유저 업데이트 영향 조회 중 에러가 발생했습니다. %d", err)
		return err
	}
	return nil
}
