package repository

import "github.com/jmoiron/sqlx"

type AuthPostgres struct {
	db *sqlx.DB
}


