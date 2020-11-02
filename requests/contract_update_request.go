package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractUpdateRequest struct {
	models.CurrencyExchangeContractMutable
	Address string `json:"address" valid:"required~ErrorBankAddressNotPassed"`
}
