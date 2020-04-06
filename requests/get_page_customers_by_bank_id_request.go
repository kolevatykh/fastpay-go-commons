package requests

type GetPageCustomersByBankIdRequest struct {
	GetPageRequest
	BankId string `json:"bankId" validate:"required"`
}
