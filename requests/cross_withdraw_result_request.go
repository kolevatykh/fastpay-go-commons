package requests

type CrossWithdrawResultRequest struct {
	WithdrawResultRequest
	CurrencyCode int `json:"currencyCode" valid:"required,min(0)"`
}
