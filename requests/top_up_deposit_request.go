package requests

type TopUpDepositRequest struct {
	SafeDealId   string `json:"safeDealId" valid:"required,uuid"`
	AddressFrom  string `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressTo    string `json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount       int64  `json:"amount" valid:"required~ErrorAmountNotPassed"`
	NeedAmount   int64  `json:"needAmount" valid:"required~ErrorAmountNotPassed"`
}
