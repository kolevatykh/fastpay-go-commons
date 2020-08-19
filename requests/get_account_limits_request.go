package requests

type GetAccountLimitsRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Address      string `json:"address" valid:"required~60302,validHex40~60301"`
}
