package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
)

type CreateCurrencyRequest struct {
	Code       int                             `json:"code" valid:"required,min(0)"`
	Name       string                          `json:"name" valid:"required"`
	Type       currency_type_enum.CurrencyType `json:"type"`
	Unit       string                          `json:"unit"`
	Symbol     string                          `json:"symbol" valid:"required,min(3),max(3)"`
	Decimals   int                             `json:"decimals" valid:"required,min(0),max(8)"`
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
