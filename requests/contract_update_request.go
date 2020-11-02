package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractUpdateRequest struct {
	models.CurrencyExchangeContractMutable
	Address string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
