package requests

type CreateInvoiceRequest struct {
	Number       string `json:"number" validate:"required,lte=255"`
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
	Amount       int64  `json:"amount" validate:"required,gt=0"`
	Description  string `json:"description"`
	Recipient    string `json:"recipient" validate:"required,validHex40"`
	Payer        string `json:"payer" validate:"omitempty,validHex40"`
}
