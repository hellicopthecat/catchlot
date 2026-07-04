package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/users/request"
)

func (db *UserRepo) RCreateUser(ctx context.Context, userReq request.CreateUserRequest) error {
	transaction, err := db.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("❌ [Transaction Failed] :: %w", err)
	}
	defer transaction.Rollback()

	insertUser, err := os.ReadFile(constants.InsertSQL + "i_users.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return fmt.Errorf("❌ [Insert User] :: %w", err)
	}

	_, err = transaction.Exec(string(insertUser), userReq.Id, userReq.Email, userReq.Social, userReq.Nickname)
	if err != nil {
		commons.BadSQLFile(err)
		return fmt.Errorf("❌ [Exec tranction User] :: %w", err)
	}

	insertUserRate, err := os.ReadFile(constants.InsertSQL + "i_users_rate.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return fmt.Errorf("❌ [Insert User Rate] :: %w", err)
	}

	transaction.Exec(string(insertUserRate), userReq.Id)

	return transaction.Commit()
}
