package reqres

import (
	"boilerplate/cmd/dso"
	"boilerplate/cmd/dto"

	"github.com/valyala/fasthttp"
)

type RegisterAdminRequest struct {
	Context *fasthttp.RequestCtx
	Item    dto.RegisterAdmin
}

type RegisterAdminResponse struct {
	Error error
}

type LoginRequest struct {
	Context *fasthttp.RequestCtx
	Issuer  string
	Item    dto.Login
}

type LoginResponse struct {
	Item         dso.Login
	AccessToken  string
	RefreshToken string
	ExpiredAt    int64
	Error        error
}

type LogoutRequest struct {
	Context *fasthttp.RequestCtx
	UserID  string
}

type LogoutResponse struct {
	Error error
}

type RefreshTokenRequest struct {
	Context      *fasthttp.RequestCtx
	RefreshToken string
	Issuer       string
}

type RefreshTokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    int64
	Error        error
}
