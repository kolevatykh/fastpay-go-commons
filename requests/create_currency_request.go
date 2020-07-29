package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/measurement-unit-enum"
)

type CreateCurrencyRequest struct {
	Code     int                                   `json:"code" validate:"required,min=0"`
	Name     string                                `json:"name" validate:"required"`
	Type     currency_type_enum.CurrencyType       `json:"type"`
	Unit     measurement_unit_enum.MeasurementUnit `json:"unit"`
	Symbol   string                                `json:"symbol" validate:"required,min=3,max=3"`
	Decimals int                                   `json:"decimals" validate:"required,min=0,max=8"`
	Enabled  bool                                  `json:"enabled"`
}
