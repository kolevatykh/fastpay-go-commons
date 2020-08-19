package requests

type GetBankBalanceTotalRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required,min(0)"`
}
