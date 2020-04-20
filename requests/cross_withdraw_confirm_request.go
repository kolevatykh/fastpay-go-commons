package requests

type CrossWithdrawConfirmRequest struct {
	WithdrawConfirmRequest
	BankId string `json:"bankId" validate:"required"`
}
