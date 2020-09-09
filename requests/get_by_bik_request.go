package requests

type GetByBikRequest struct {
	Bik string `json:"bik" valid:"required"`
}
