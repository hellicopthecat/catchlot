package repo

import "database/sql"

type OfficialLottoRepo struct {
	db *sql.DB
}

func InitOfficialLottoRepo(db *sql.DB) *OfficialLottoRepo {
	return &OfficialLottoRepo{db: db}
}
