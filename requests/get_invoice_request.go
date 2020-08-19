package requests

type GetInvoiceRequest struct {
	Number    string `json:"number" valid:"required~60352,maxstringlength(255)"`
	Recipient string `json:"recipient" valid:"required~60302,validHex40~60301"`
}
