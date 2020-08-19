package requests

type GetListByPayerRequest struct {
	Payer string `json:"payer" valid:"required~60302,validHex40~60301"`
}
