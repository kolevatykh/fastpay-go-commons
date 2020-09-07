package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type TopUpDepositRequest struct {
	SafeDealId   string                  `json:"safeDealId" valid:"required,uuid"`
	AddressFrom  string                  `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	AddressTo    string                  `json:"addressTo" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyInfo models.CurrencyDealInfo `json:"currencyInfo" valid:"required"`
	Amount       int64                   `json:"amount" valid:"required~ErrorAmountNotPassed"`
	NeedAmount   int64                   `json:"needAmount" valid:"required~ErrorAmountNotPassed"`
}
