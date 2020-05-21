package requests

type GetPageByPayerRequest struct {
	GetPageRequest
	Payer string `json:"payer" validate:"required,validHex40"`
}
