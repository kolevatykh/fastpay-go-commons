package requests

type CrossWithdrawConfirmRequest struct {
	WithdrawConfirmRequest
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
