package service

import (
	"boilerplate/cmd/dso"
	"boilerplate/cmd/repository"
	repoReqres "boilerplate/cmd/repository/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"errors"
)

type AuthorService interface {
	GetAuthor()
	DetailAuthor()
	CreateAuthor(*serviceReqres.CreateAuthorRequest) serviceReqres.CreateAuthorResponse
	UpdateAuthor()
	DeleteAuthor()
}

type authorService struct {
	dao repository.DAO
}

func NewAuthorService(dao repository.DAO) AuthorService {
	return &authorService{dao: dao}
}

func (a *authorService) GetAuthor() {

}

func (a *authorService) DetailAuthor() {

}

func (a *authorService) CreateAuthor(serviceReq *serviceReqres.CreateAuthorRequest) (serviceRes serviceReqres.CreateAuthorResponse) {
	repoReq := repoReqres.CreateAuthorRequest{
		Context: serviceReq.Context,
		Item: dso.Author{
			Name:        serviceReq.Item.Name,
			Address:     serviceReq.Item.Address,
			PhoneNumber: serviceReq.Item.PhoneNumber,
		},
	}
	repoRes := a.dao.NewAuthorQuery().CreateAuthor(&repoReq)
	if repoRes.Error != nil {
		serviceRes.Error = errors.New("internal server error")
		return
	}
	return
}

func (a *authorService) UpdateAuthor() {

}

func (a *authorService) DeleteAuthor() {

}
