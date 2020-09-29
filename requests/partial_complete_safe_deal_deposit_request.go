package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type PartialCompleteSafeDealDepositRequest struct {
	SafeDealId   string                  `json:"safeDealId" valid:"required,uuid"`
	AddressTo    string                  `json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount       int64                   `json:"amount" valid:"range(0|9223372036854775807)~ErrorAmountNegative"`
	CurrencyInfo models.CurrencyDealInfo `json:"currencyInfo" valid:"required"`
}
