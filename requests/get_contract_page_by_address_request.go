package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-type-enum"
)

type GetContractPageByBankRequest struct {
	GetContractPageRequest
	Address string `json:"address" valid:"required~ErrorBankAddressNotPassed"`
}

func (getContractPage *GetContractPageByBankRequest) SetDefaults() {
	if getContractPage.Type == currency_exchange_contracts_type_enum.Undefined {
		getContractPage.Type = currency_exchange_contracts_type_enum.CurrencyExchangeContract
	}
}
