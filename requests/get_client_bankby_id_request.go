package requests

type GetClientBankByIdRequest struct {
	BankId string `json:"bankId" validate:"required"`
}
