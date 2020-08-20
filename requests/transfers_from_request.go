package requests

type TransfersFromRequest struct {
	AddressFrom   string            `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            []TransfersFromTo `json:"to" valid:"required"`
	CurrencyCode  int               `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	TransactionId string            `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}

type TransfersFromTo struct {
	To      string `json:"to" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	Amount  int64  `json:"amount" valid:"required~ErrorAmountNotPassed"`
	Payload string `json:"payload"`
}
