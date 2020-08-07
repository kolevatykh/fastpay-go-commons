package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"

type GetListCurrencyRequest struct {
	Type currency_type_enum.CurrencyType `json:"type"`
}

func (getListCurrencyRequest *GetListCurrencyRequest) SetDefaults() {
	if getListCurrencyRequest.Type == currency_type_enum.Undefined {
		getListCurrencyRequest.Type = currency_type_enum.DigitalCurrency
	}
}
