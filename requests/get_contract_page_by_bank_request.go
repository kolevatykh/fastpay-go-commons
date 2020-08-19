package requests

type GetContractPageByBankRequest struct {
	BankId string `json:"bankId" valid:"required~60364"`
	GetContractPageRequest
}
