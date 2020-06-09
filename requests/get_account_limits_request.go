package requests

type GetAccountLimitsRequest struct {
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
	Address      string `json:"address" validate:"required,validHex40"`
}
