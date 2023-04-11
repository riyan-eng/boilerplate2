package service

import "boilerplate/cmd/repository"

type AuthenticationService interface {
	RegisterAdmin()
}

type authenticationService struct {
	dao repository.DAO
}

func NewAuthenticationService(dao repository.DAO) AuthenticationService {
	return &authenticationService{
		dao: dao,
	}
}

func (a *authenticationService) RegisterAdmin() {

}
