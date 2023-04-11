package reqres

import (
	"boilerplate/cmd/dso"

	"github.com/valyala/fasthttp"
)

type CreateAuthorRequest struct {
	Context *fasthttp.RequestCtx
	Item    dso.Author
}

type CreateAuthorResponse struct {
	Error error
}
