package requests

type ExecuteClearingRequest struct {
	CurrencyCode int `json:"currencyCode" validate:"required,min=0"`
}
