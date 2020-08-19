package requests

type GetByCodeRequest struct {
	Code int `json:"code" valid:"required,min(0)"`
}
