package requests

type GetInvoiceRequest struct {
	Number    string `json:"number" valid:"required~ErrorNumberNotPassed,maxstringlength(255)"`
	Recipient string `json:"recipient" valid:"required~ErrorAddressNotPassed,validHex40~ErrorAddressNotFollowingRegex"`
}
