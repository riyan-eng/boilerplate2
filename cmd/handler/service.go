package handler

import "boilerplate/cmd/service"

type MicroServiceServer struct {
	authenticationService service.AuthenticationService
}

func NewMicroService(authenticationService service.AuthenticationService) *MicroServiceServer {
	return &MicroServiceServer{
		authenticationService: authenticationService,
	}
}
