package requests

type SetIssueLimitRequest struct {
	Address      string `json:"address" valid:"required,validHex40"`
	Value        int64  `json:"value" valid:"gte(0)"`
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
}
