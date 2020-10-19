package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ExecuteContractRequest struct {
	ContractInfo  models.CurrencyContractRoutingItem `json:"contractInfo" valid:"required"`
	AddressFrom   string                             `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            string                             `json:"to" valid:"validHex40or64~ErrorAddressNotFollowingRegex"`
	Amount        int64                              `json:"amount" valid:"required~ErrorAmountNotPassed"`
	TransactionId string                             `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
	Payload       string                             `json:"payload"`
}
