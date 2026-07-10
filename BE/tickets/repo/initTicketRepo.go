package repo

import (
	"database/sql"

	"github.com/hellicopthecat/catchlot/gakSoo/repo"
)

type TicketRepo struct {
	db  *sql.DB
	gak *repo.GakSooCache
}

func InitTicketRepo(db *sql.DB, gak *repo.GakSooCache) *TicketRepo {
	return &TicketRepo{
		db:  db,
		gak: gak,
	}
}
