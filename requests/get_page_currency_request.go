package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"

type GetPageCurrencyRequest struct {
	GetPageRequest
	Type currency_type_enum.CurrencyType `json:"type"`
}

func (getPageCurrencyRequest *GetPageCurrencyRequest) SetDefaults() {
	if getPageCurrencyRequest.Type == currency_type_enum.Undefined {
		getPageCurrencyRequest.Type = currency_type_enum.DigitalCurrency
	}
}
