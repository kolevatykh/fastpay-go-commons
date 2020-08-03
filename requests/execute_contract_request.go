package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ExecuteContractRequest struct {
	ContractInfo  models.CurrencyContractRoutingItem `json:"contractInfo" validate:"required,dive"`
	AddressFrom   string                             `json:"addressFrom" validate:"required,validHex40"`
	To            string                             `json:"to" validate:"omitempty,validHex40or64"`
	Amount        int64                              `json:"amount" validate:"required"`
	TransactionId string                             `json:"transactionId" validate:"required,uuid4"`
	Payload       string                             `json:"payload"`
}
