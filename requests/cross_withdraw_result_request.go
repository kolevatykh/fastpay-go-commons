package requests

type CrossWithdrawResultRequest struct {
	WithdrawResultRequest
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	BankAddress  string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
