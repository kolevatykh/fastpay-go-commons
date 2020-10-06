package requests

import "github.com/SolarLabRU/fastpay-go-commons/models"

type ContractAwardedRequest struct {
	MakeDepositRequest
	IsLastDepositNoNeedTransfers bool                    `json:"isLastDepositNoNeedTransfers"`
	CurrencyInfo                 models.CurrencyDealInfo `json:"currencyInfo"`
}
