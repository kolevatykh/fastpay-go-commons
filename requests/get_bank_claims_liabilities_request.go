package requests

type GetBankClaimsLiabilitiesRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Bank         string `json:"bank" valid:"required~ErrorBankAddressNotPassed,validHex40"`
}
