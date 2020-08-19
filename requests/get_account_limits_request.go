package requests

type GetAccountLimitsRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	Address      string `json:"address" valid:"required,validHex40"`
}
