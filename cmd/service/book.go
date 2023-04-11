package service

import "boilerplate/cmd/repository"

type BookService interface {
	GetBook()
	DetailBook()
	CreateBook()
	UpdateBook()
	DeleteBook()
}

type bookService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BookService {
	return &bookService{dao: dao}
}

func (a *bookService) GetBook() {

}

func (a *bookService) DetailBook() {

}

func (a *bookService) CreateBook() {

}

func (a *bookService) UpdateBook() {

}

func (a *bookService) DeleteBook() {

}
