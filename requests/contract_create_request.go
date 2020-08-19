package requests

type ContractCreateRequest struct {
	Id                   string  `json:"id" valid:"required~60304"`
	AddressAccountSell   string  `json:"addressAccountSell" valid:"required~60302,validHex40~60301"`
	AddressAccountBuy    string  `json:"addressAccountBuy" valid:"required~60302,validHex40~60301"`
	AddressCommission    string  `json:"addressCommission" valid:"required~60302,validHex40~60301"`
	CurrencyCodeSell     int     `json:"currencyCodeSell" valid:"required~60318,range(0|999)~60317"`
	CurrencyCodeBuy      int     `json:"currencyCodeBuy" valid:"required~60318,range(0|999)~60317"`
	CurrencySymbolSell   string  `json:"currencySymbolSell" valid:"required~60337,stringlength(3|3)"`
	CurrencySymbolBuy    string  `json:"currencySymbolBuy" valid:"required~60337,stringlength(3|3)"`
	Price                float64 `json:"price" valid:"required~60313,range(0.0000000001|9223372036854775807)"`
	FractionalCommission float64 `json:"fractionalCommission" valid:"range(0|1)"`
	MaxCommission        int64   `json:"maxCommission" valid:"range(0|9223372036854775807)"`
	MinAmount            int64   `json:"minAmount" valid:"range(0|9223372036854775807)~60312"`
	MaxAmount            int64   `json:"maxAmount" valid:"required,range(0|9223372036854775807)~60312"`
	StartDate            int64   `json:"startDate" valid:"required~60332,range(0|9223372036854775807)~60333"`
	EndDate              int64   `json:"endDate" valid:"required~60332,range(0|9223372036854775807)~60333"`
}
