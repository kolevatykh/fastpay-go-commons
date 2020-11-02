package requests

type WithdrawConfirmRequest struct {
	TechnicalSignRequest
	AddressFrom   string `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	TxId          string `json:"txId" valid:"required~ErrorTxIdSNotPassed"`
	CurrencyCode  int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	TransactionId string `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}

type WithdrawRejectRequest WithdrawConfirmRequest
