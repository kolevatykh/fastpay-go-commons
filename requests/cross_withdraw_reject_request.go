package requests

type CrossWithdrawRejectRequest struct {
	WithdrawRejectRequest
	BankId string `json:"bankId" validate:"required"`
}
