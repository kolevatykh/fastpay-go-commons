package requests

type GetAllClaimsRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
