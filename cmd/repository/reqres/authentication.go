package reqres

import (
	"boilerplate/cmd/dso"

	"github.com/valyala/fasthttp"
)

type RegisterAdminRequest struct {
	Context *fasthttp.RequestCtx
	Item    dso.RegisterAdmin
}

type RegisterAdminResponse struct {
	Error error
	Item  dso.RegisterAdmin
}

type LoginRequest struct {
	Item    dso.Login
	Context *fasthttp.RequestCtx
}

type LoginResponse struct {
	Item  dso.Login
	Error error
}
