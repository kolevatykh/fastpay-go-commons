package requests

type GetAccountBalanceRequest struct {
	Address      string `json:"address" validate:"required,validHex40"`
	CurrencyCode int    `json:"currencyCode" validate:"required,min=0"`
}
