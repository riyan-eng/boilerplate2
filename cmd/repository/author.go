package repository

import (
	"boilerplate/cmd/repository/reqres"
	"database/sql"
	"fmt"
)

type AuthorQuery interface {
	GetAuthor()
	DetailAuthor()
	CreateAuthor(*reqres.CreateAuthorRequest) reqres.CreateAuthorResponse
	UpdateAuthor()
	DeleteAuthor()
}

type authorQuery struct {
	database *sql.DB
}

func (a *authorQuery) GetAuthor() {

}

func (a *authorQuery) DetailAuthor() {

}

func (a *authorQuery) CreateAuthor(repoReq *reqres.CreateAuthorRequest) (repoRes reqres.CreateAuthorResponse) {
	q := fmt.Sprintf(`
		insert into author(name, address, phone_number) values ('%v', '%v', '%v')
	`, repoReq.Item.Name, repoReq.Item.Address, repoReq.Item.PhoneNumber)
	fmt.Println(q)
	return
}

func (a *authorQuery) UpdateAuthor() {

}

func (a *authorQuery) DeleteAuthor() {

}
