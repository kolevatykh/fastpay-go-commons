package requests

type CrossWithdrawRejectRequest struct {
	WithdrawRejectRequest
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
