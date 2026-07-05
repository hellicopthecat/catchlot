package repo

import (
	"database/sql"
	"os"

	"github.com/hellicopthecat/catchlot/commons"
	"github.com/hellicopthecat/catchlot/constants"
)

func (r UserRepo) RFindUserByEmail(email string) (string, error) {
	q, err := os.ReadFile(constants.SelectSQL + "users/s_findByUserByEmail.sql")
	if err != nil {
		commons.BadSQLFile(err)
		return "", err
	}
	var userEmail string
	scanErr := r.db.QueryRow(string(q), email).Scan(&userEmail)
	if scanErr == sql.ErrNoRows {
		return "", nil
	}
	if scanErr != nil {
		return "", scanErr
	}
	return userEmail, nil
}
