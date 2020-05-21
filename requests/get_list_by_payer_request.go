package requests

type GetListByPayerRequest struct {
	Payer string `json:"payer" validate:"required,validHex40"`
}
