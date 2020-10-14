package requests

type GetPageCrossTransactionHistoryRequest struct {
	GetPageRequest
	BankId string `json:"bankId" valid:"required~ErrorBankIdNotPassed"`
}
