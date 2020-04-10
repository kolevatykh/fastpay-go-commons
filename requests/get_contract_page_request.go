package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/filter-contract-state"

type GetContractPageRequest struct {
	PageSize     int32                                     `json:"pageSize" validate:"required"`
	Bookmark     string                                    `json:"bookmark"`
	EndDate      int64                                     `json:"endDate" validate:"omitempty,min=0"`
	Status       filter_contract_state.FilterContractState `json:"status" validate:"omitempty,min=0"`
	Address      string                                    `json:"address" validate:"omitempty,validHex40"`
	CurrencyCode int                                       `json:"currencyCode" validate:"omitempty,min=0"`
}
