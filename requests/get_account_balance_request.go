package requests

type GetAccountBalanceRequest struct {
	Address      string `json:"address" valid:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" valid:"required,min(0)"`
}
