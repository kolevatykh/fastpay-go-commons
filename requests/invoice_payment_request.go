package requests

type InvoicePaymentRequest struct {
	Number       string `json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Recipient    string `json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Payer        string `json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount       int64  `json:"amount" valid:"required~ErrorAmountNotPassed,range(1|9223372036854775807)~ErrorAmountNegative"`
}
