package requests

type RefundRequest struct {
	TechnicalSignRequest
	AddressFrom   string `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode  int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Payload       string `json:"payload"`
	TransactionId string `json:"transactionId" valid:"required~ErrorTransactionIdNotPassed,uuidv4"`
}
