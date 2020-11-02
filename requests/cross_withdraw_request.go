package requests

type CrossWithdrawRequest struct {
	WithdrawRequest
	BankAddress string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
