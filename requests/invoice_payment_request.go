package requests

type InvoicePaymentRequest struct {
	Number       string `json:"number" valid:"required,lte(255)"`
	Recipient    string `json:"recipient" valid:"required,validHex40"`
	Payer        string `json:"payer" valid:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	Amount       int64  `json:"amount" valid:"required,gte(1)"`
}
