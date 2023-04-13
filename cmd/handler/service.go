package handler

import "boilerplate/cmd/service"

type MicroServiceServer struct {
	authorService         service.AuthorService
	bookService           service.BookService
	authenticationService service.AuthenticationService
}

func NewMicroService(authorService service.AuthorService, bookService service.BookService, authenticationService service.AuthenticationService) *MicroServiceServer {
	return &MicroServiceServer{
		authorService:         authorService,
		bookService:           bookService,
		authenticationService: authenticationService,
	}
}
