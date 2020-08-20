package requests

type GetByCodeRequest struct {
	Code int `json:"code" valid:"required,range(0|9223372036854775807)"`
}
