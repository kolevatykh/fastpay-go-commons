package requests

type GetInvoiceRequest struct {
	Number    string `json:"number" valid:"required,lte(255)"`
	Recipient string `json:"recipient" valid:"required,validHex40"`
}
