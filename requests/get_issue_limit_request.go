package requests

type GetIssueLimitRequest struct {
	Address      string `json:"address" validate:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
}
