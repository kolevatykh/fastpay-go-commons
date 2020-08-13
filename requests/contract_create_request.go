package requests

type ContractCreateRequest struct {
	Id                   string  `json:"id" validate:"required"`
	AddressAccountSell   string  `json:"addressAccountSell" validate:"required,validHex40"`
	AddressAccountBuy    string  `json:"addressAccountBuy" validate:"required,validHex40"`
	AddressCommission    string  `json:"addressCommission" validate:"required,validHex40"`
	CurrencyCodeSell     int     `json:"currencyCodeSell" validate:"required,min=0"`
	CurrencyCodeBuy      int     `json:"currencyCodeBuy" validate:"required,min=0"`
	CurrencySymbolSell   string  `json:"currencySymbolSell" validate:"required,min=3,max=3"`
	CurrencySymbolBuy    string  `json:"currencySymbolBuy" validate:"required,min=3,max=3"`
	CurrencyUnitSell     string  `json:"currencyUnitSell"`
	CurrencyUnitBuy      string  `json:"currencyUnitBuy"`
	Price                float64 `json:"price" validate:"required,gte=0.0000000001"`
	FractionalCommission float64 `json:"fractionalCommission" validate:"gte=0,lte=1"`
	MaxCommission        int64   `json:"maxCommission" validate:"gte=0"`
	MinAmount            int64   `json:"minAmount" validate:"omitempty,min=0"`
	MaxAmount            int64   `json:"maxAmount" validate:"required,min=0"`
	StartDate            int64   `json:"startDate" validate:"required,min=0"`
	EndDate              int64   `json:"endDate" validate:"required,min=0"`
}
