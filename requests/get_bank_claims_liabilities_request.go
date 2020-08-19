package requests

type GetBankClaimsLiabilitiesRequest struct {
	CurrencyCode int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	Bank         string `json:"bank" valid:"required~60364,validHex40~60367"`
}
