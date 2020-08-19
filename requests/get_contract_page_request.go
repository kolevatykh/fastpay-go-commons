package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/filter-contract-state-enum"
)

type GetContractPageRequest struct {
	PageSize     int32                                          `json:"pageSize" valid:"required"`
	Bookmark     string                                         `json:"bookmark"`
	EndDate      int64                                          `json:"endDate" valid:"min(0)"`
	Status       filter_contract_state_enum.FilterContractState `json:"status" valid:"min(0)"`
	Address      string                                         `json:"address" valid:"validHex40"`
	CurrencyCode int                                            `json:"currencyCode" valid:"min(0)"`
}
