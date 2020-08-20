package requests

import "github.com/SolarLabRU/fastpay-go-commons/enums/invoice-state-enum"

type UpdateInvoiceRequest struct {
	Number       string                          `json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)~ErrorNumberInvoiceNotFolowingRegex"`
	Recipient    string                          `json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	Payer        string                          `json:"payer" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
	State        invoice_state_enum.InvoiceState `json:"state" valid:"required~ErrorInvoiceStateNotPassed"`
	CurrencyCode int                             `json:"currencyCode" valid:"required~ErrorCurrencyCodeNotPassed,range(0|999)~ErrorCurrencyCodeRange"`
	ErrorCode    int                             `json:"errorCode"`
}
