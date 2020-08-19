package requests

type ExecuteClearingRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required,min(0)"`
}
