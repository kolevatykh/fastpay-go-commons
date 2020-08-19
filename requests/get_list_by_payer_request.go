package requests

type GetListByPayerRequest struct {
	Payer string `json:"payer" valid:"required,validHex40"`
}
