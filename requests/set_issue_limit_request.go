package requests

type SetIssueLimitRequest struct {
	TechnicalSignRequest
	Address      string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Value        int64  `json:"value" valid:"range(0|9223372036854775807)"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
