package requests

type GetBankBalanceTotalRequest struct {
	CurrencyCode int `json:"currencyCode" validate:"required,min=0"`
}
