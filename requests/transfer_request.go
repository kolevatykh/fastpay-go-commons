package requests

type TransferRequest struct {
	AddressFrom   string  `json:"addressFrom" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	To            string  `json:"to" valid:"required~ErrorAddressNotPassed,validHex40or64~ErrorAddressOrIdentifierNotFolowingRegex"`
	CurrencyCode  int     `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount        int64   `json:"amount" valid:"required~ErrorAmountNotPassed"`
	Payload       string  `json:"payload"`
	MsgHash       string  `json:"msgHash"`
	Sig           SignDto `json:"sig"`
	InvoiceNumber string  `json:"invoiceNumber" valid:"maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Exp           int64   `json:"exp"`
	TransactionId string  `json:"transactionId" valid:"uuidv4"`
}
