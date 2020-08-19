package requests

type GetBankClaimsLiabilitiesRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required,gte(0),lte(999)"`
	Bank         string `json:"bank" valid:"required,validHex40"`
}
