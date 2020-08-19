package requests

type SetPaymentErrorRequest struct {
	Number       string `json:"number" valid:"required,lte(255)"`
	Recipient    string `json:"recipient" valid:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	ErrorCode    int    `json:"errorCode" valid:"required"`
}
