package requests

type WithdrawConfirmRequest struct {
	AddressFrom   string `json:"addressFrom" valid:"required,validHex40"`
	TxId          string `json:"txId" valid:"required"`
	CurrencyCode  int    `json:"currencyCode" valid:"required,min(0)"`
	TransactionId string `json:"transactionId" valid:"required,uuidv4"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
