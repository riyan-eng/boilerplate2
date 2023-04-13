package repository

import "database/sql"

type DAO interface {
	NewAuthorQuery() AuthorQuery
	NewBookQuery() BookQuery
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

func (d *dao) NewAuthorQuery() AuthorQuery {
	return &authorQuery{
		database: d.dbpostgre,
	}
}

func (d *dao) NewBookQuery() BookQuery {
	return &bookQuery{}
}

func (d *dao) NewAuthenticationQuery() AuthenticationQuery {
	return &authenticationQuery{
		database: d.dbpostgre,
	}
}
