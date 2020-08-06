package requests

type GetContractByIdRequest struct {
	Id string `json:"id" validate:"required"`
}
