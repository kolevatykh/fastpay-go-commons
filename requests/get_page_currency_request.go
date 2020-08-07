package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/currency-type-enum"

type GetPageCurrencyRequest struct {
	PageSize int32                           `json:"pageSize" validate:"required"`
	Bookmark string                          `json:"bookmark"`
	Type     currency_type_enum.CurrencyType `json:"type"`
}
