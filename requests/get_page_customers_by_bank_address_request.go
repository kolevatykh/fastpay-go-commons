package requests

type GetPageCustomersByBankAddressRequest struct {
	GetPageRequest
	Address string `json:"bankAddress" valid:"required~ErrorBankAddressNotPassed"`
}
