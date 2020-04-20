package requests

type CrossWithdrawRequest struct {
	WithdrawRequest
	BankId string `json:"bankId" validate:"required"`
}
