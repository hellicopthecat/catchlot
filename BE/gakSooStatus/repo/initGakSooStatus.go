package repo

import "database/sql"

type GakSooStatusRepo struct {
	db *sql.DB
}

func InitGakSooStatus(db *sql.DB) *GakSooStatusRepo {
	return &GakSooStatusRepo{
		db: db,
	}
}
