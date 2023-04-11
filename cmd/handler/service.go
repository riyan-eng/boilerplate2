package handler

import "boilerplate/cmd/service"

type MicroServiceServer struct {
	authorService service.AuthorService
	bookService   service.BookService
}

func NewMicroService(authorService service.AuthorService, bookService service.BookService) *MicroServiceServer {
	return &MicroServiceServer{
		authorService: authorService,
		bookService:   bookService,
	}
}
