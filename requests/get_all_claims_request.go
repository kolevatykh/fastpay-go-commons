package requests

type GetAllClaimsRequest struct {
	CurrencyCode int `json:"currencyCode" validate:"required,gte=0,lte=999"`
}
