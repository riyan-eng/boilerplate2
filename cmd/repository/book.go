package repository

type BookQuery interface {
	GetBook()
	DetailBook()
	CreateBook()
	UpdateBook()
	DeleteBook()
}

type bookQuery struct{}

func (b *bookQuery) GetBook() {

}

func (b *bookQuery) DetailBook() {

}

func (b *bookQuery) CreateBook() {

}

func (b *bookQuery) UpdateBook() {

}

func (b *bookQuery) DeleteBook() {

}
