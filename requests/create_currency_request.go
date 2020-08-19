package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
)

type CreateCurrencyRequest struct {
	Code       int                             `json:"code" valid:"required~60318,range(0|999)~60317"`
	Name       string                          `json:"name" valid:"required"`
	Type       currency_type_enum.CurrencyType `json:"type"`
	Unit       string                          `json:"unit"`
	Symbol     string                          `json:"symbol" valid:"required~60337,stringlength(3|3)"`
	Decimals   int                             `json:"decimals" valid:"required~60338,range(0|8)~60339"`
	Properties map[string]string               `json:"properties"`
	Enabled    bool                            `json:"enabled"`
}

func (createCurrencyRequest *CreateCurrencyRequest) SetDefaults() {
	if createCurrencyRequest.Type == currency_type_enum.Undefined {
		createCurrencyRequest.Type = currency_type_enum.DigitalCurrency
	}
	if createCurrencyRequest.Properties == nil {
		createCurrencyRequest.Properties = make(map[string]string)
	}
}
