package requests

type CreateCurrencyRequest struct {
	Code     int    `json:"code" validate:"required,min=0"`
	Name     string `json:"name" validate:"required"`
	Symbol   string `json:"symbol" validate:"required,min=3,max=3"`
	Decimals int    `json:"decimals" validate:"required,min=0,max=8"`
	Enabled  bool   `json:"enabled"`
}
