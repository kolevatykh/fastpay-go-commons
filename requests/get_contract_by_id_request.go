package requests

type GetContractByIdRequest struct {
	Id     string `json:"id" valid:"required"`
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
