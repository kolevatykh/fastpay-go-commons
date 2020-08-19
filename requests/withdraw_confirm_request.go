package requests

type WithdrawConfirmRequest struct {
	AddressFrom   string `json:"addressFrom" valid:"required~60302,validHex40~60301"`
	TxId          string `json:"txId" valid:"required~60366"`
	CurrencyCode  int    `json:"currencyCode" valid:"required~60318,range(0|999)~60317"`
	TransactionId string `json:"transactionId" valid:"required~60366,uuidv4"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
