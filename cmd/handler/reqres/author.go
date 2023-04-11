package reqres

type CreateAuthorRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type CreateAuthorResponse struct {
	Id string `json:"id"`
}
