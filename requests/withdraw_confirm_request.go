package requests

type WithdrawConfirmRequest struct {
	AddressFrom   string `json:"addressFrom" validate:"required,validHex40"`
	TxId          string `json:"txId" validate:"required"`
	CurrencyCode  int    `json:"currencyCode" validate:"required,min=0"`
	TransactionId string `json:"transactionId" validate:"required,uuid4"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
