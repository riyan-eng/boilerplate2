package reqres

import (
	"boilerplate/cmd/dto"

	"github.com/valyala/fasthttp"
)

type GetAuthorRequest struct {
	Context *fasthttp.RequestCtx
	Page    int
	Limit   int
	Search  string
}

type GetAuthorResponse struct {
	Error error
}

type CreateAuthorRequest struct {
	Context *fasthttp.RequestCtx
	Item    dto.Author
}

type CreateAuthorResponse struct {
	Error error
}
