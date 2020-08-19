package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ExecuteContractRequest struct {
	ContractInfo  models.CurrencyContractRoutingItem `json:"contractInfo" valid:"required"`
	AddressFrom   string                             `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	To            string                             `json:"to" valid:"validHex40or64~60301"`
	Amount        int64                              `json:"amount" valid:"required~60313"`
	TransactionId string                             `json:"transactionId" valid:"required~60331,uuidv4"`
	Payload       string                             `json:"payload"`
}
