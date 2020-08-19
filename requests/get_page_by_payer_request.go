package requests

type GetPageByPayerRequest struct {
	GetPageRequest
	Payer string `json:"payer" valid:"required,validHex40"`
}
