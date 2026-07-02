package repo

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
	"github.com/hellicopthecat/catchlot/users/request"
)

func CreateUser(ctx context.Context, db *sql.Conn, userReq request.CreateUserRequest) error {
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("❌ [Transaction Failed] :: %d", err)
	}
	defer transaction.Rollback()

	insertUser, err := os.ReadFile(constants.InsertSQL + "i_users.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return fmt.Errorf("❌ [Insert User] :: %d", err)
	}

	transaction.Exec(string(insertUser), userReq.Id, userReq.Email, userReq.Social, userReq.Nickname)

	insertUserRate, err := os.ReadFile(constants.InsertSQL + "i_users_rate.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return fmt.Errorf("❌ [Insert User Rate] :: %d", err)
	}

	transaction.Exec(string(insertUserRate), userReq.Id)

	return transaction.Commit()
}
