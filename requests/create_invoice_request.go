package requests

type CreateInvoiceRequest struct {
	Number       string `json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	Amount       int64  `json:"amount" valid:"required~ErrorAmountNotPassed,range(1|9223372036854775807)"`
	Description  string `json:"description"`
	Recipient    string `json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Payer        string `json:"payer" valid:"validHex40~ErrorAddressNotFollowingRegex"`
}
