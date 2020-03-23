package requests

type GetByCodeRequest struct {
	Code int `json:"code" validate:"required,min=0"`
}
