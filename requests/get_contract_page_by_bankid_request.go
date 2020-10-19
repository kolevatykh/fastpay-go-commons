package requests

import (
	"github.com/SolarLabRU/fastpay-go-commons/enums/currency-exchange-contracts-type-enum"
)

type GetContractPageByBankRequest struct {
	GetContractPageRequest
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}

func (getContractPage *GetContractPageByBankRequest) SetDefaults() {
	if getContractPage.Type == currency_exchange_contracts_type_enum.Undefined {
		getContractPage.Type = currency_exchange_contracts_type_enum.CurrencyExchangeContract
	}
}
