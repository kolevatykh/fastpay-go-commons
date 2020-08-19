package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/filter-contract-state-enum"
)

type GetContractPageRequest struct {
	PageSize     int32                                          `json:"pageSize" valid:"required"`
	Bookmark     string                                         `json:"bookmark"`
	EndDate      int64                                          `json:"endDate" valid:"range(0|9223372036854775807)"`
	Status       filter_contract_state_enum.FilterContractState `json:"status" valid:"range(0|9223372036854775807)"`
	Address      string                                         `json:"address" valid:"validHex40~60301"`
	CurrencyCode int                                            `json:"currencyCode" valid:"range(0|999)~60317"`
}
