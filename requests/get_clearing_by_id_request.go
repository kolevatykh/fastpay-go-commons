package requests

type GetClearingByIdRequest struct {
	Id string `json:"id" validate:"required"`
}
