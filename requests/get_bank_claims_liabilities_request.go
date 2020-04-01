package requests

type GetBankClaimsLiabilitiesRequest struct {
	CurrencyCode int    `json:"currencyCode" validate:"required,gte=0,lte=999"`
	Bank         string `json:"bank" validate:"required,validHex40"`
}
