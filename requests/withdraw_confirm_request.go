package requests

type WithdrawConfirmRequest struct {
	AddressFrom   string `json:"addressFrom" validate:"required,validHex40"`
	TxId          string `json:"txId" validate:"required"` // TODO
	CurrencyCode  int    `json:"currencyCode" validate:"required,min=0"`
	TransactionId int    `json:"transactionId" validate:"required,min=0"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
