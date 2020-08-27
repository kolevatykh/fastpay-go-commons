package requests

type ContractAwardedRequest struct {
	MakeDepositRequest
	IsLastDepositNoNeedTransfers bool `json:"isLastDepositNoNeedTransfers"`
	CurrencyCode                 int  `json:"currencyCode"`
}
