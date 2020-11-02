package requests

type GetPageCustomersByAddressRequest struct {
	GetPageRequest
	Address string `json:"address" valid:"required~ErrorBankAddressNotPassed"`
}
