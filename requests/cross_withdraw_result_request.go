package requests

type CrossWithdrawResultRequest struct {
	WithdrawResultRequest
	CurrencyCode int `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
}
