package requests

type GetClientBankByAddressRequest struct {
	Address string `json:"address" valid:"required~ErrorBankAddressNotPassed"`
}
