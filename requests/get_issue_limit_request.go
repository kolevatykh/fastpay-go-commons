package requests

type GetIssueLimitRequest struct {
	Address      string `json:"address" valid:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
}
