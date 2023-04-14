package service

import (
	"boilerplate/cmd/dso"
	"boilerplate/cmd/repository"
	repoReqres "boilerplate/cmd/repository/reqres"
	serviceReqres "boilerplate/cmd/service/reqres"
	"boilerplate/config"
	"boilerplate/internal/util"
	"boilerplate/pkg"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type AuthenticationService interface {
	RegisterAdmin(*serviceReqres.RegisterAdminRequest) serviceReqres.RegisterAdminResponse
	Login(*serviceReqres.LoginRequest) serviceReqres.LoginResponse
	Logout(*serviceReqres.LogoutRequest) serviceReqres.LogoutResponse
	RefreshToken(*serviceReqres.RefreshTokenRequest) serviceReqres.RefreshTokenResponse
}

type authenticationService struct {
	dao repository.DAO
}

func NewAuthenticationService(dao repository.DAO) AuthenticationService {
	return &authenticationService{
		dao: dao,
	}
}

func (a *authenticationService) RegisterAdmin(serviceReq *serviceReqres.RegisterAdminRequest) (serviceRes serviceReqres.RegisterAdminResponse) {
	repoReq := repoReqres.RegisterAdminRequest{
		Context: serviceReq.Context,
		Item: dso.RegisterAdmin{
			UserName:     serviceReq.Item.UserName,
			Password:     util.GenerateHash(serviceReq.Item.Password),
			Email:        serviceReq.Item.Email,
			UserTypeCode: "admin",
			Pin:          util.GenerateHash(serviceReq.Item.Password),
			PhoneNumber:  serviceReq.Item.PhoneNumber,
			UserDataID:   "",
			CreatedBy:    "",
		},
	}
	repoRes := a.dao.NewAuthenticationQuery().RegisterAdmin(&repoReq)
	if repoRes.Error != nil {
		log.Println(repoRes.Error.Error())
		serviceRes.Error = repoRes.Error
		return
	}
	return
}

func (a *authenticationService) Login(serviceReq *serviceReqres.LoginRequest) (serviceRes serviceReqres.LoginResponse) {
	repoReq := repoReqres.LoginRequest{
		Context: serviceReq.Context,
		Item: dso.Login{
			UserName: serviceReq.Item.UserName,
		},
	}
	repoRes := a.dao.NewAuthenticationQuery().Login(&repoReq)
	if repoRes.Error == sql.ErrNoRows {
		log.Println(repoRes.Error.Error())
		serviceRes.Error = errors.New("username doesn't exist")
		return
	}
	if !util.VerifyHash(repoRes.Item.Password, serviceReq.Item.Password) {
		serviceRes.Error = errors.New("password didn't match")
		return
	}
	accessToken, refreshToken, expiredAt, err := util.GenerateJWT(serviceReq.Issuer, repoRes.Item.ID, "", 45)
	if err != nil {
		log.Println(err.Error())
		serviceRes.Error = errors.New("error generate token")
		return
	}
	serviceRes.AccessToken = accessToken
	serviceRes.RefreshToken = refreshToken
	serviceRes.ExpiredAt = expiredAt
	serviceRes.Item = repoRes.Item
	return
}

func (a *authenticationService) Logout(serviceReq *serviceReqres.LogoutRequest) (serviceRes serviceReqres.LogoutResponse) {
	err := config.Redis.Del(serviceReq.Context, fmt.Sprintf("token-%s", serviceReq.UserID)).Err()
	if err != nil {
		log.Println(err.Error())
		serviceRes.Error = errors.New("internal server error")
		return
	}
	return
}

func (a *authenticationService) RefreshToken(serviceReq *serviceReqres.RefreshTokenRequest) (serviceRes serviceReqres.RefreshTokenResponse) {
	claims, err := util.ParseToken(serviceReq.RefreshToken, "AllYourBaseRefresh")
	if err != nil {
		log.Println(err.Error())
		serviceRes.Error = errors.New(pkg.ERROR_REQUEST)
		return
	}
	if err := util.ValidateToken(claims, true, serviceReq.Context); err != nil {
		log.Println(err.Error())
		serviceRes.Error = errors.New(pkg.ERROR_REQUEST)
		return
	}
	accessToken, refreshToken, expiredAt, err := util.GenerateJWT(serviceReq.Issuer, claims.UserID, "", 45)
	if err != nil {
		log.Println(err.Error())
		serviceRes.Error = errors.New("error generate token")
		return
	}
	serviceRes.AccessToken = accessToken
	serviceRes.RefreshToken = refreshToken
	serviceRes.ExpiredAt = expiredAt
	return
}
