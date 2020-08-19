package requests

type ContractCreateRequest struct {
	Id                   string  `json:"id" valid:"required~60304"`
	AddressAccountSell   string  `json:"addressAccountSell" valid:"required~60302,validHex40~60301"`
	AddressAccountBuy    string  `json:"addressAccountBuy" valid:"required~60302,validHex40~60301"`
	AddressCommission    string  `json:"addressCommission" valid:"required~60302,validHex40~60301"`
	CurrencyCodeSell     int     `json:"currencyCodeSell" valid:"required~60318,min(0)~60317"`
	CurrencyCodeBuy      int     `json:"currencyCodeBuy" valid:"required~60318,min(0)~60317"`
	CurrencySymbolSell   string  `json:"currencySymbolSell" valid:"required~60337,min(3),max(3)"`
	CurrencySymbolBuy    string  `json:"currencySymbolBuy" valid:"required~60337,min(3),max(3)"`
	Price                float64 `json:"price" valid:"required~60313,gte(0.0000000001)"`
	FractionalCommission float64 `json:"fractionalCommission" valid:"gte(0),lte(1)"`
	MaxCommission        int64   `json:"maxCommission" valid:"gte(0)"`
	MinAmount            int64   `json:"minAmount" valid:"min(0)~60312"`
	MaxAmount            int64   `json:"maxAmount" valid:"required,min(0)~60312"`
	StartDate            int64   `json:"startDate" valid:"required~60332,min(0)~60333"`
	EndDate              int64   `json:"endDate" valid:"required~60332,min(0)~60333"`
}
