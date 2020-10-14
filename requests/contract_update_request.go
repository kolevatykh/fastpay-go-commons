package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractUpdateRequest struct {
	models.CurrencyExchangeContractMutable
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
