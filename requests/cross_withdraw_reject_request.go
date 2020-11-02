package requests

type CrossWithdrawRejectRequest struct {
	WithdrawRejectRequest
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
