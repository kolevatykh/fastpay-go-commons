package requests

type CreateInvoiceRequest struct {
	Number       string `json:"number" valid:"required,lte(255)"`
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	Amount       int64  `json:"amount" valid:"required,gte(1)"`
	Description  string `json:"description"`
	Recipient    string `json:"recipient" valid:"required,validHex40"`
	Payer        string `json:"payer" valid:"validHex40"`
}
