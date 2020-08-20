package requests

type ExecuteClearingRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
