package requests

type GetAllClaimsRequest struct {
	CurrencyCode int `json:"currencyCode" valid:"required,gte(0),lte(999)"`
}
