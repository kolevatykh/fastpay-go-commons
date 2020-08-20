package requests

type GetAccountLimitsRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Address      string `json:"address" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
