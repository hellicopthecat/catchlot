package repo

import "database/sql"

type UserRepo struct {
	db *sql.DB
}

func InitUserRepo(db *sql.DB) UserRepo {
	return UserRepo{db: db}
}
