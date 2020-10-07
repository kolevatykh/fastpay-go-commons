package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-type-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/filter-contract-state-enum"
	"github.com/SolarLabRU/fastpay-go-commons/enums/operation-deal-type-enum"
)

type GetContractPageRequest struct {
	PageSize      int32                                                               `json:"pageSize" valid:"required"`
	Bookmark      string                                                              `json:"bookmark"`
	EndDate       int64                                                               `json:"endDate" valid:"range(0|9223372036854775807)"`
	Status        filter_contract_state_enum.FilterContractState                      `json:"status" valid:"range(0|9223372036854775807)"`
	OperationType operation_deal_type_enum.OperationDealType                          `json:"operationType" valid:"range(0|9223372036854775807)"`
	Address       string                                                              `json:"address" valid:"validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode  int                                                                 `json:"currencyCode" valid:"range(0|999)~ErrorCurrencyCodeRange"`
	Type          currency_exchange_contracts_type_enum.CurrencyExchangeContractsType `json:"type"`
}

func (getContractPage *GetContractPageRequest) SetDefaults() {
	if getContractPage.Type == currency_exchange_contracts_type_enum.Undefined {
		getContractPage.Type = currency_exchange_contracts_type_enum.CurrencyExchangeContract
	}
}
