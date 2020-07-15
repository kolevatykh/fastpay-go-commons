package requests

type SetPaymentErrorRequest struct {
	Number       string `json:"number" validate:"required,lte=255"`
	Recipient    string `json:"recipient" validate:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
	ErrorCode    int    `json:"errorCode" validate:"required"`
}
