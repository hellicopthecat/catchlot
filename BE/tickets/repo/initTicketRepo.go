package repo

import "database/sql"

type TicketRepo struct {
	db *sql.DB
}

func InitTicketRepo(db *sql.DB) *TicketRepo {
	return &TicketRepo{
		db: db,
	}
}
