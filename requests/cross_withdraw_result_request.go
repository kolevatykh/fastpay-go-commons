package requests

type CrossWithdrawResultRequest struct {
	WithdrawResultRequest
	CurrencyCode int `json:"currencyCode" validate:"required,min=0"`
}
