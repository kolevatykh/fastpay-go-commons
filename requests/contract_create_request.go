package requests

type ContractCreateRequest struct {
	Id                   string  `json:"id" valid:"required"`
	AddressAccountSell   string  `json:"addressAccountSell" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressAccountBuy    string  `json:"addressAccountBuy" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressCommission    string  `json:"addressCommission" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCodeSell     int     `json:"currencyCodeSell" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencyCodeBuy      int     `json:"currencyCodeBuy" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	CurrencySymbolSell   string  `json:"currencySymbolSell" valid:"required~ErrorSymbolNotPassed,stringlength(3|3)"`
	CurrencySymbolBuy    string  `json:"currencySymbolBuy" valid:"required~ErrorSymbolNotPassed,stringlength(3|3)"`
	Price                float64 `json:"price" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)"`
	FractionalCommission float64 `json:"fractionalCommission" valid:"range(0|1)"`
	MaxCommission        int64   `json:"maxCommission" valid:"range(0|9223372036854775807)"`
	MinAmount            int64   `json:"minAmount" valid:"range(0|9223372036854775807)~ErrorAmountNegative"`
	MaxAmount            int64   `json:"maxAmount" valid:"required~ErrorAmountNotPassed,range(0|9223372036854775807)~ErrorAmountNegative"`
	StartDate            int64   `json:"startDate" valid:"required~ErrorTimestampNotPassed,range(0|9223372036854775807)~ErrorTimestampNegative"`
	EndDate              int64   `json:"endDate" valid:"required~ErrorTimestampNotPassed,range(0|9223372036854775807)~ErrorTimestampNegative"`
}
