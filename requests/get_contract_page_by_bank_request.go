package requests

type GetContractPageByBankRequest struct {
	BankId string `json:"bankId" validate:"required"`
	GetContractPageRequest
}
