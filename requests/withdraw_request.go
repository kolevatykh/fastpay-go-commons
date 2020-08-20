package requests

type WithdrawRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Amount        int64   `json:"amount" valid:"required~ErrorAmountNotPassed"`
	CurrencyCode  int     `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Payload       string  `json:"payload"`
	TransactionId string  `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
	MsgHash       string  `json:"msgHash" valid:"required"`
	Exp           int64   `json:"exp" valid:"required~ErrorTimestampNotPassed"`
	Sig           SignDto `json:"sig" valid:"required"`
}
