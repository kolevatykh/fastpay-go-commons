package requests

type GetPageCustomersByBankIdRequest struct {
	GetPageRequest
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
