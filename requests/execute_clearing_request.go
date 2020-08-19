package requests

type ExecuteClearingRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
}
