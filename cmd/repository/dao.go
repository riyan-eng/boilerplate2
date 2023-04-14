package repository

import "database/sql"

type DAO interface {
	NewAuthenticationQuery() AuthenticationQuery
}

type dao struct {
	dbpostgre *sql.DB
}

func NewDao(dbpostgre *sql.DB) DAO {
	return &dao{
		dbpostgre: dbpostgre,
	}
}

func (d *dao) NewAuthenticationQuery() AuthenticationQuery {
	return &authenticationQuery{
		database: d.dbpostgre,
	}
}
