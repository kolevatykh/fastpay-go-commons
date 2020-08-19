package requests

type GetPageByPayerRequest struct {
	GetPageRequest
	Payer string `json:"payer" valid:"required~60302,validHex40~60301"`
}
