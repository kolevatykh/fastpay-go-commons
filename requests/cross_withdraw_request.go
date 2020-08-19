package requests

type CrossWithdrawRequest struct {
	WithdrawRequest
	BankId string `json:"bankId" valid:"required"`
}
