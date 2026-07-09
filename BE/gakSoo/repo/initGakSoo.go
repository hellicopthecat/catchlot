package repo

import "database/sql"

type GakSooRepo struct {
	db *sql.DB
}

func InitGakSoo(db *sql.DB) *GakSooRepo {
	return &GakSooRepo{
		db: db,
	}
}
