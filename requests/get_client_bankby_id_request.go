package requests

type GetClientBankByIdRequest struct {
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
