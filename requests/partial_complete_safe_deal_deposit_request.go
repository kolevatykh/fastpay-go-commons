package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type PartialCompleteSafeDealDepositRequest struct {
	SafeDealId   string                  `json:"safeDealId" valid:"required,uuid"`
	AddressTo    string                  `json:"addressAcceptor" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount       int64                   `json:"amount"`
	CurrencyInfo models.CurrencyDealInfo `json:"currencyInfo" valid:"required"`
}