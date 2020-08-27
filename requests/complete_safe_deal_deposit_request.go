package requests

type CompleteSafeDealDepositRequest struct {
	SafeDealId   string `json:"safeDealId" valid:"required,uuid"`
	AddressTo    string `json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
}
