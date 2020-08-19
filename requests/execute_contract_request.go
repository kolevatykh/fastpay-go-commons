package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ExecuteContractRequest struct {
	ContractInfo  models.CurrencyContractRoutingItem `json:"contractInfo" valid:"required"`
	AddressFrom   string                             `json:"addressFrom" valid:"required,validHex40"`
	To            string                             `json:"to" valid:"validHex40or64"`
	Amount        int64                              `json:"amount" valid:"required"`
	TransactionId string                             `json:"transactionId" valid:"required,uuidv4"`
	Payload       string                             `json:"payload"`
}
