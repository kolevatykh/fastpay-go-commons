package requests

type SetIssueLimitRequest struct {
	Address      string `json:"address" valid:"required~60302,validHex40~60301"`
	Value        int64  `json:"value" valid:"range(0|9223372036854775807)"`
	CurrencyCode int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
}
