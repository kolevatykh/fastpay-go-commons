package requests

type SetPaymentErrorRequest struct {
	Number       string `json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Recipient    string `json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	CurrencyCode int    `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	ErrorCode    int    `json:"errorCode" valid:"required"`
}
