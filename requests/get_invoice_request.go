package requests

type GetInvoiceRequest struct {
	Number    string `json:"number" validate:"required,lte=255"`
	Recipient string `json:"recipient" validate:"required,validHex40"`
}
