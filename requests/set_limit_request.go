package requests

type SetLimitRequest struct {
	GetLimitRequest
	CurrencyCode int   `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Value        int64 `json:"value" valid:"range(0|9223372036854775807)"`
}
